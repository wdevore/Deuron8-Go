package graphs

import (
	"fmt"
	"image/color"

	"github.com/inkyblackness/imgui-go/v2"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

// This graph renders chains of Spikes
// Each spike is a vertical lines about N pixels in height
// Each row is seperated by ~2px.
// Poisson spikes are orange, AP spikes are green.
// Poisson is drawn first then AP.
//
// Graph is shaped like this:
//      .----------------> +X
//  1   :  |   ||     |   | |       ||     |
//  2   :    |   |   ||     ||     |    |        <-- a row ~2px height
//  3   :   |    |    |         | |   |     |
//      v
//      +Y
//
// Only the X-axis is mapped Y is simply a height is graph-space.
//
// This graph also shows the Neuron's Post spike (i.e. the output of the neuron)

const (
	maxVerticalBarsLimit = 250
	spikeRowOffset       = 2
	spikeHeight          = 10
	cellSpikeHeight      = 30
	cellLineThickness    = 2.0
	panelHeight          = 300
)

type spikeGraph struct {
	// Vertical time bar markers
	showMarkers     bool
	timePos         int
	showPoissonData bool
	showStimData    bool

	spikeColor                imgui.PackedColor
	stimulusColor             imgui.PackedColor
	noiseColor                imgui.PackedColor
	verticalMarkerLightColor  imgui.PackedColor
	verticalCursorMarkerColor imgui.PackedColor

	toolTipBackgroundColor imgui.PackedColor
	toolTipForgroundColor  imgui.PackedColor
	toolTipMinCorner       imgui.Vec2
	toolTipMaxCorner       imgui.Vec2

	p1 imgui.Vec2
	p2 imgui.Vec2
}

// NewSpikeGraph creates imgui graph
func NewSpikeGraph() api.IGraph {
	o := new(spikeGraph)
	o.showPoissonData = true
	o.showStimData = true

	o.toolTipBackgroundColor = imgui.Packed(color.Gray{Y: 32})
	o.toolTipForgroundColor = imgui.Packed(color.Gray{Y: 128})
	o.verticalMarkerLightColor = imgui.Packed(color.Gray{Y: 64})
	o.verticalCursorMarkerColor = imgui.Packed(color.Gray{Y: 128})

	o.stimulusColor = imgui.Packed(color.RGBA{R: 166, G: 255, B: 77, A: 255})
	o.noiseColor = imgui.Packed(color.RGBA{R: 255, G: 255, B: 0, A: 255})
	o.spikeColor = imgui.Packed(color.White)

	o.toolTipMinCorner = imgui.Vec2{}
	o.toolTipMaxCorner = imgui.Vec2{}

	o.p1 = imgui.Vec2{}
	o.p2 = imgui.Vec2{}

	return o
}

func (g *spikeGraph) Draw(environment api.IEnvironment, vertPos int) {
	imgui.SetNextWindowPosV(imgui.Vec2{X: 0.0, Y: float32(vertPos)}, imgui.ConditionOnce, imgui.Vec2{})
	config := environment.Config()

	moData, _ := config.Data().(*model.ConfigJSON)
	imgui.SetNextWindowSizeV(imgui.Vec2{X: float32(moData.WindowWidth - 10), Y: float32(panelHeight + 20)}, imgui.ConditionAlways)

	imgui.Begin("Spike Graph")

	g.drawHeader(environment)

	imgui.Separator()

	g.drawGraph(environment)

	imgui.End()
}

func (g *spikeGraph) drawHeader(environment api.IEnvironment) {
	if imgui.TreeNode("Controls##1") {
		config := environment.Config()
		imgui.PushItemWidth(200)

		moData, _ := config.Data().(*model.ConfigJSON)
		rangeStart := int32(moData.RangeStart)
		rangeEnd := int32(moData.RangeEnd)
		duration := int32(moData.Duration)

		changedS := imgui.DragIntV("RangeStart##1", &rangeStart, 1.0, 0, int32(moData.RangeEnd), "%d")

		imgui.SameLine()

		changedE := imgui.DragIntV("RangeEnd##1", &rangeEnd, 1.0, rangeStart, duration, "%d")

		if changedS || changedE {
			if rangeStart < rangeEnd {
				config.Changed()
				moData.RangeStart = int(rangeStart)
				moData.RangeEnd = int(rangeEnd)
			}
		}

		scrollVelocity := float32(moData.Scroll)

		changed := imgui.SliderFloatV("Scroll Velocity", &scrollVelocity, -5.0, 5.0, "%.2f", 1.0)
		if changed {
			moData.Scroll = float64(scrollVelocity)
		}

		velocity := ScrollVelocity(moData.Scroll)
		rangeDx := rangeEnd - rangeStart

		if moData.Scroll < 0 {
			rangeStart += int32(velocity)
			// Left
			if rangeStart > 0 {
				rangeEnd = rangeStart + rangeDx
			} else {
				rangeStart = 0
				rangeEnd = rangeStart + rangeDx
			}
			config.Changed()
			moData.RangeStart = int(rangeStart)
			moData.RangeEnd = int(rangeEnd)
		} else if moData.Scroll > 0 {
			rangeEnd += int32(velocity)
			if rangeEnd < duration {
				rangeStart = rangeEnd - rangeDx
			} else {
				rangeEnd = duration
				rangeStart = rangeEnd - rangeDx
			}
			config.Changed()
			moData.RangeStart = int(rangeStart)
			moData.RangeEnd = int(rangeEnd)
		}

		// If above slider is released we clear the velocity.
		if !imgui.IsItemActive() {
			moData.Scroll = 0.0
		}

		imgui.SameLineV(250, 100)

		imgui.PushItemWidth(400)
		imgui.Checkbox("Show Noise", &g.showPoissonData)
		imgui.PopItemWidth()

		imgui.SameLineV(350, 100)
		imgui.Checkbox("Show Stimulus", &g.showStimData)

		barRange := moData.RangeEnd - moData.RangeStart
		if barRange < maxVerticalBarsLimit {
			imgui.SameLineV(500, 100)

			// Limit bars to less than Max because Drawlist is limited to 2^16 items.
			imgui.Checkbox("Show Markers", &g.showMarkers)
		} else {
			g.showMarkers = false
		}

		imgui.PopItemWidth()

		imgui.TreePop()
	}
}

func (g *spikeGraph) drawGraph(environment api.IEnvironment) {
	drawList := imgui.WindowDrawList()
	canvasPos := imgui.CursorScreenPos()
	canvasSize := imgui.ContentRegionAvail()

	cx := canvasSize.X
	cy := canvasSize.Y

	if cx < 50.0 {
		cx = 50.0
	}
	if cy < 50.0 {
		cy = 50.0
	}

	canvasSize.X = cx
	canvasSize.Y = cy

	g.toolTipMinCorner.X = canvasPos.X
	g.toolTipMinCorner.Y = canvasPos.Y

	g.toolTipMaxCorner.X = canvasPos.X + canvasSize.X
	g.toolTipMaxCorner.Y = canvasPos.Y + canvasSize.Y

	drawList.AddRectFilled(g.toolTipMinCorner, g.toolTipMaxCorner, g.toolTipBackgroundColor)

	if g.showMarkers {
		if imgui.IsWindowHovered() {
			imgui.BeginTooltip()
			imgui.Text(fmt.Sprintf("%d", g.timePos))
			imgui.EndTooltip()
		}

		// Below doesn't work in Inky's binding don't include clipping
		// A visible button scaled to the size of the canvas is used for hover checking
		// imgui.InvisibleButtonV("canvas", canvasSize)
		// if imgui.IsItemHovered() {
		// 	imgui.BeginTooltip()
		// 	imgui.Text(fmt.Sprintf("%d", g.timePos))
		// 	imgui.EndTooltip()
		// }

		drawList.AddRect(g.toolTipMinCorner, g.toolTipMaxCorner, g.toolTipForgroundColor)

		g.drawVerticalMarkers(environment.Config(), drawList)
	}

	if g.showPoissonData {
		g.drawNoise(environment, drawList)
	}

	if g.showStimData {
		g.drawData(environment, drawList)
	}

	g.drawSomaData(environment, drawList)
}

// -----------------------------------------------------------------------
// Only a window of data is shown based on the RangeStart and End
// The range needs to be mapped to the graph window via lerping.
// Scrolling adjusts the Range by moving both the Start and End.
// -----------------------------------------------------------------------

func (g *spikeGraph) drawVerticalMarkers(config api.IModel, drawList imgui.DrawList) {
	mousePosX := imgui.MousePos().X

	// Mapped data coords
	uX := 0.0
	wX := 0.0
	plvx := 0.0

	moData, _ := config.Data().(*model.ConfigJSON)
	rangeStart := int32(moData.RangeStart)
	rangeEnd := int32(moData.RangeEnd)

	// timePos tracks the actual time regardless of scrolling so it always
	// starts at the current range start value.
	timePos := int(rangeStart)
	rangeDx := rangeEnd - rangeStart
	canvasSize := imgui.ContentRegionAvail()
	canvasPos := imgui.CursorScreenPos()

	// "t" is a counter over the range "size". timePos is the actual
	// time value capture for the tooltip
	for t := int32(0); t < rangeDx; t++ {
		// We want the markers to track with time as well, so we map "t".
		uX = MapSampleToUnit(float64(t), 0.0, float64(rangeDx))
		wX = MapUnitToWindow(uX, 0.0, float64(canvasSize.X))

		lX, lY := MapWindowToLocal(wX, 0.0, canvasPos)

		if float64(mousePosX) > plvx && float64(mousePosX) < lX {
			// Draw cursor bar instead of time marker
			g.p1.X = float32(lX)
			g.p1.Y = float32(lY)
			g.p2.X = float32(lX)
			g.p2.Y = float32(lY) + canvasSize.Y
			drawList.AddLine(g.p1, g.p2, g.verticalCursorMarkerColor)
			g.timePos = timePos // + 1
		} else {
			// Draw time marker
			g.p1.X = float32(lX)
			g.p1.Y = float32(lY)
			g.p2.X = float32(lX)
			g.p2.Y = float32(lY) + canvasSize.Y
			drawList.AddLine(g.p1, g.p2, g.verticalMarkerLightColor)
		}

		plvx = lX // Capture previous value for interval testing
		timePos++
	}
}

func (g *spikeGraph) drawSomaData(environment api.IEnvironment, drawList imgui.DrawList) {
	wY := 200.0 + 40.0 // Offset from any previous rows
	config := environment.Config()
	moData, _ := config.Data().(*model.ConfigJSON)
	rangeStart := int32(moData.RangeStart)
	rangeEnd := int32(moData.RangeEnd)
	canvasSize := imgui.ContentRegionAvail()

	samples := environment.Samples()
	canvasPos := imgui.CursorScreenPos()

	somaData := samples.SomaData()

	if somaData != nil && len(somaData) > 0 {
		for t := rangeStart; t < rangeEnd; t++ {
			// Draw channel
			if somaData[t].Output() == 1 { // A spike = 1
				// The sample value needs to be mapped
				uX := MapSampleToUnit(float64(t), float64(rangeStart), float64(rangeEnd))
				wX := MapUnitToWindow(uX, 0.0, float64(canvasSize.X))
				lX, lY := MapWindowToLocal(wX, wY, canvasPos)

				g.p1.X = float32(lX)
				g.p1.Y = float32(lY)
				g.p2.X = float32(lX)
				g.p2.Y = float32(lY) + spikeHeight
				drawList.AddLine(g.p1, g.p2, g.spikeColor)
			}
		}
	}
}

func (g *spikeGraph) drawData(environment api.IEnvironment, drawList imgui.DrawList) {
	wY := 100.0 + 20.0 // Offset from any previous rows
	config := environment.Config()
	moData, _ := config.Data().(*model.ConfigJSON)
	rangeStart := int32(moData.RangeStart)
	rangeEnd := int32(moData.RangeEnd)
	canvasSize := imgui.ContentRegionAvail()

	samples := environment.Samples()
	canvasPos := imgui.CursorScreenPos()

	synapseData := samples.SynapticData()

	for i, channel := range synapseData {
		if channel != nil {
			if i < 10 { // The Stimulus data is the lower channels
				for t := rangeStart; t < rangeEnd; t++ {
					// Draw channel
					if channel[t].Input() == 1 { // A spike = 1
						// The sample value needs to be mapped
						uX := MapSampleToUnit(float64(t), float64(rangeStart), float64(rangeEnd))
						wX := MapUnitToWindow(uX, 0.0, float64(canvasSize.X))
						lX, lY := MapWindowToLocal(wX, wY, canvasPos)

						g.p1.X = float32(lX)
						g.p1.Y = float32(lY)
						g.p2.X = float32(lX)
						g.p2.Y = float32(lY) + spikeHeight
						drawList.AddLine(g.p1, g.p2, g.stimulusColor)
					}
				}
				// Update row/y value and offset by a few pixels
				wY += spikeHeight + spikeRowOffset
			}
		}
	}
}

func (g *spikeGraph) drawNoise(environment api.IEnvironment, drawList imgui.DrawList) {
	wY := 0.0 // Offset from border. 0 is underneath it.
	config := environment.Config()
	moData, _ := config.Data().(*model.ConfigJSON)
	rangeStart := int32(moData.RangeStart)
	rangeEnd := int32(moData.RangeEnd)
	canvasSize := imgui.ContentRegionAvail()

	samples := environment.Samples()
	canvasPos := imgui.CursorScreenPos()

	synapseData := samples.SynapticData()

	for i, channel := range synapseData {
		if channel != nil {
			if i > 9 {
				for t := rangeStart; t < rangeEnd; t++ {
					// Draw channel
					if channel[t].Input() == 1 { // A spike = 1
						// The sample value needs to be mapped
						uX := MapSampleToUnit(float64(t), float64(rangeStart), float64(rangeEnd))
						wX := MapUnitToWindow(uX, 0.0, float64(canvasSize.X))
						lX, lY := MapWindowToLocal(wX, wY, canvasPos)

						g.p1.X = float32(lX)
						g.p1.Y = float32(lY)
						g.p2.X = float32(lX)
						g.p2.Y = float32(lY) + spikeHeight
						drawList.AddLine(g.p1, g.p2, g.noiseColor)
					}
				}
				// Update row/y value and offset by a few pixels
				wY += spikeHeight + spikeRowOffset
			}
		}
	}
}
