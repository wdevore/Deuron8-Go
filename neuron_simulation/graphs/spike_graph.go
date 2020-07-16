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
)

type spikeGraph struct {
	// Vertical time bar markers
	showMarkers     bool
	timePos         int
	showPoissonData bool
	showStimData    bool

	toolTipBackgroundColor imgui.PackedColor
	toolTipForgroundColor  imgui.PackedColor
	toolTipMinCorner       imgui.Vec2
	toolTipMaxCorner       imgui.Vec2
}

// NewSpikeGraph creates imgui graph
func NewSpikeGraph() api.IGraph {
	o := new(spikeGraph)
	o.showPoissonData = true
	o.showStimData = true
	// o.toolTipBackgroundColor = imgui.Packed(color.RGBA{R: 64, G: 64, B: 64, A: 255})
	o.toolTipBackgroundColor = imgui.Packed(color.Gray{Y: 32})
	o.toolTipForgroundColor = imgui.Packed(color.Gray{Y: 128})
	o.toolTipMinCorner = imgui.Vec2{}
	o.toolTipMaxCorner = imgui.Vec2{}
	return o
}

func (g *spikeGraph) Draw(config, sim api.IModel, samples api.ISamples, vertPos int) {
	imgui.SetNextWindowPosV(imgui.Vec2{X: 0.0, Y: float32(vertPos)}, imgui.ConditionOnce, imgui.Vec2{})

	moData, _ := config.Data().(*model.ConfigJSON)
	imgui.SetNextWindowSizeV(imgui.Vec2{X: float32(moData.WindowWidth - 10), Y: float32(200 + 20)}, imgui.ConditionAlways)

	imgui.Begin("Spike Graph")

	g.drawHeader(config, sim)

	imgui.Separator()

	g.drawGraph(config, sim, samples)

	imgui.End()
}

func (g *spikeGraph) drawHeader(config, sim api.IModel) {
	if imgui.TreeNode("Controls##1") {
		imgui.PushItemWidth(200)

		moData, _ := config.Data().(*model.ConfigJSON)
		rangeStart := int32(moData.RangeStart)
		rangeEnd := int32(moData.RangeEnd)
		duration := int32(moData.Duration)

		changedS := imgui.DragIntV("RangeStart##1", &rangeStart, 1.0, 1, int32(moData.RangeEnd), "%d")

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
			config.Changed()
			moData.Scroll = float64(scrollVelocity)
		}

		imgui.SameLineV(450, 30)

		imgui.PushItemWidth(300)
		imgui.Checkbox("Show Noise", &g.showPoissonData)
		imgui.PopItemWidth()
		imgui.SameLineV(350, 30)
		imgui.Checkbox("Show Stimulus", &g.showStimData)

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

func (g *spikeGraph) drawGraph(config, sim api.IModel, samples api.ISamples) {
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

	// A visible button scaled to the size of the canvas is used for hover checking
	imgui.InvisibleButtonV("canvas", canvasSize)

	drawList.AddRectFilled(g.toolTipMinCorner, g.toolTipMaxCorner, g.toolTipBackgroundColor)
	if imgui.IsItemHovered() {
		imgui.BeginTooltip()
		imgui.Text(fmt.Sprintf("%d", g.timePos))
		imgui.EndTooltip()
	}

	drawList.AddRect(g.toolTipMinCorner, g.toolTipMaxCorner, g.toolTipForgroundColor)

	g.drawData(config, sim, samples)
}

func (g *spikeGraph) drawData(config, sim api.IModel, samples api.ISamples) {

}