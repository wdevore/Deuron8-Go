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
	thresholdColor           imgui.PackedColor
	zeroColor                imgui.PackedColor
	verticalMarkerLightColor imgui.PackedColor

	p1 imgui.Vec2
	p2 imgui.Vec2
}

// NewSomaPspGraph creates imgui graph
func NewSomaPspGraph() api.IGraph {
	o := new(somaPspGraph)

	o.lineColor = imgui.Packed(color.RGBA{R: 255, G: 127, B: 0, A: 255})
	o.verticalMarkerLightColor = imgui.Packed(color.Gray{Y: 64})
	o.thresholdColor = imgui.Packed(color.RGBA{R: 127, G: 127, B: 255, A: 255})
	o.zeroColor = imgui.Packed(color.RGBA{R: 200, G: 200, B: 127, A: 255})

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

		barRange := moData.RangeEnd - moData.RangeStart
		if barRange < maxVerticalBarsLimit {

			// Limit bars to less than Max because Drawlist is limited to 2^16 items.
			imgui.Checkbox("Show Markers", &g.showMarkers)
		} else {
			g.showMarkers = false
		}

		imgui.PopItemWidth()

		imgui.TreePop()
	}
}

func (g *somaPspGraph) drawGraph(environment api.IEnvironment) {
	drawList := imgui.WindowDrawList()

	if g.showMarkers {
		g.drawVerticalMarkers(environment.Config(), drawList)
	}

	g.drawData(environment, drawList)
	g.drawHorizontalLines(environment, drawList)
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

func (g *somaPspGraph) drawHorizontalLines(environment api.IEnvironment, drawList imgui.DrawList) {
	simMod := environment.Sim()
	moData, _ := simMod.Data().(*model.SimJSON)
	canvasSize := imgui.ContentRegionAvail()
	canvasPos := imgui.CursorScreenPos()

	samples := environment.Samples()
	somaData := samples.SomaData()

	if len(somaData) > 0 {
		// ----------------------------------------------------------------
		// Threshold line
		// ----------------------------------------------------------------
		// The sample value needs to be mapped
		uY := MapSampleToUnit(moData.Neuron.Threshold, samples.SomaPspMin(), samples.SomaPspMax())
		// graph space has +Y downward, but the data is oriented as +Y upward
		// so we flip in unit-space.
		uY = 1.0 - uY

		wY := MapUnitToWindow(uY, 0.0, float64(canvasSize.Y))

		_, lY := MapWindowToLocal(0.0, wY, canvasPos)

		wBx := MapUnitToWindow(0.0, 0.0, float64(canvasSize.X))
		wEx := MapUnitToWindow(1.0, 0.0, float64(canvasSize.X))
		lBx, _ := MapWindowToLocal(wBx, 0.0, canvasPos)
		lEx, _ := MapWindowToLocal(wEx, 0.0, canvasPos)

		g.p1.X = float32(lBx)
		g.p1.Y = float32(lY)
		g.p2.X = float32(lEx)
		g.p2.Y = float32(lY)
		drawList.AddLine(g.p1, g.p2, g.thresholdColor)

		// ----------------------------------------------------------------
		// Zero line
		// ----------------------------------------------------------------
		uY = MapSampleToUnit(0.0, samples.SomaPspMin(), samples.SomaPspMax())
		uY = 1.0 - uY
		wY = MapUnitToWindow(uY, 0.0, float64(canvasSize.Y))
		_, lY = MapWindowToLocal(0.0, wY, canvasPos)

		g.p1.X = float32(lBx)
		g.p1.Y = float32(lY)
		g.p2.X = float32(lEx)
		g.p2.Y = float32(lY)
		drawList.AddLine(g.p1, g.p2, g.zeroColor)
	}
}
