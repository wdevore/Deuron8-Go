package graphs

import (
	"image/color"

	"github.com/inkyblackness/imgui-go/v2"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

// This graph renders a soma's psp values

type somaPspGraph struct {
	// Vertical time bar markers
	showMarkers bool

	timePos int

	lineColor                imgui.PackedColor
	verticalMarkerLightColor imgui.PackedColor

	p1 imgui.Vec2
	p2 imgui.Vec2
}

// NewSynapseWeightsGraph creates imgui graph
func NewSomaPspGraph() api.IGraph {
	o := new(somaPspGraph)

	o.lineColor = imgui.Packed(color.RGBA{R: 255, G: 127, B: 0, A: 255})
	o.verticalMarkerLightColor = imgui.Packed(color.Gray{Y: 64})

	o.p1 = imgui.Vec2{}
	o.p2 = imgui.Vec2{}

	return o
}

func (g *somaPspGraph) Draw(environment api.IEnvironment, vertPos int) {
	imgui.SetNextWindowPosV(imgui.Vec2{X: 0.0, Y: float32(vertPos)}, imgui.ConditionOnce, imgui.Vec2{})
	config := environment.Config()

	moData, _ := config.Data().(*model.ConfigJSON)
	imgui.SetNextWindowSizeV(imgui.Vec2{X: float32(moData.WindowWidth - 10), Y: float32(PspPanelHeight)}, imgui.ConditionAlways)

	imgui.Begin("Soma Psp Graph")

	g.drawHeader(environment)

	imgui.Separator()

	if !environment.IsRunning() {
		g.drawGraph(environment)
	}

	imgui.End()
}

func (g *somaPspGraph) drawHeader(environment api.IEnvironment) {
	if imgui.TreeNode("Controls##5") {
		config := environment.Config()
		imgui.PushItemWidth(200)

		moData, _ := config.Data().(*model.ConfigJSON)
		rangeStart := int32(moData.RangeStart)
		rangeEnd := int32(moData.RangeEnd)
		duration := int32(moData.Duration)

		changedS := imgui.DragIntV("RangeStart##5", &rangeStart, 1.0, 0, int32(moData.RangeEnd), "%d")

		imgui.SameLine()

		changedE := imgui.DragIntV("RangeEnd##5", &rangeEnd, 1.0, rangeStart, duration, "%d")

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

		imgui.TreePop()
	}
}

func (g *somaPspGraph) drawGraph(environment api.IEnvironment) {
	drawList := imgui.WindowDrawList()

	if g.showMarkers {
		g.drawVerticalMarkers(environment.Config(), drawList)
	}

	g.drawData(environment, drawList)
}

// -----------------------------------------------------------------------
// Only a window of data is shown based on the RangeStart and End
// The range needs to be mapped to the graph window via lerping.
// Scrolling adjusts the Range by moving both the Start and End.
// -----------------------------------------------------------------------

func (g *somaPspGraph) drawVerticalMarkers(config api.IModel, drawList imgui.DrawList) {
	// Mapped data coords
	uX := 0.0
	wX := 0.0

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

		// Draw time marker
		g.p1.X = float32(lX)
		g.p1.Y = float32(lY)
		g.p2.X = float32(lX)
		g.p2.Y = float32(lY) + canvasSize.Y
		drawList.AddLine(g.p1, g.p2, g.verticalMarkerLightColor)

		timePos++
	}
}

// Draw a curve graph for the selected synapse
func (g *somaPspGraph) drawData(environment api.IEnvironment, drawList imgui.DrawList) {
	config := environment.Config()
	moData, _ := config.Data().(*model.ConfigJSON)
	rangeStart := int32(moData.RangeStart)
	rangeEnd := int32(moData.RangeEnd)
	canvasSize := imgui.ContentRegionAvail()

	samples := environment.Samples()
	canvasPos := imgui.CursorScreenPos()

	somaData := samples.SomaData()

	sY := 0.0
	plY := 0.0 // previously mapped y value
	plX := 0.0 // previously mapped x value

	for t := rangeStart; t < rangeEnd; t++ {
		if len(somaData) > 0 {
			sY = somaData[t].Psp()

			// The sample value needs to be mapped
			uX := MapSampleToUnit(float64(t), float64(rangeStart), float64(rangeEnd))
			uY := MapSampleToUnit(sY, samples.SomaPspMin(), samples.SomaPspMax())

			wX := MapUnitToWindow(uX, 0.0, float64(canvasSize.X))

			// graph space has +Y downward, but the data is oriented as +Y upward
			// so we flip in unit-space.
			uY = 1.0 - uY
			wY := MapUnitToWindow(uY, 0.0, float64(canvasSize.Y))

			lX, lY := MapWindowToLocal(wX, wY, canvasPos)

			g.p1.X = float32(plX)
			g.p1.Y = float32(plY)
			g.p2.X = float32(lX)
			g.p2.Y = float32(lY)
			drawList.AddLine(g.p1, g.p2, g.lineColor)

			// if model.bug println("vt: ", vt) end
			plX = lX
			plY = lY
		}
	}
}
