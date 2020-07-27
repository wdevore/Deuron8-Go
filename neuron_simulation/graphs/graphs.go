package graphs

import (
	"math"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/misc"
)

const (
	SpikePanelHeight   = 300
	SurgePanelHeight   = 200
	PspPanelHeight     = 200
	WeightsPanelHeight = 200

	SomaPspPanelHeight    = 200
	SomaAPFastPanelHeight = 200
	SomaAPSlowPanelHeight = 200
)

var (
	p1 = imgui.Vec2{}
	p2 = imgui.Vec2{}
)

// MapSampleToUnit from sample-space to unit-space where unit-space is 0->1
func MapSampleToUnit(v, min, max float64) float64 {
	return misc.Linear(min, max, v)
}

// MapUnitToWindow from unit-space to window-space
func MapUnitToWindow(v, min, max float64) float64 {
	return misc.Lerp(min, max, v)
}

// MapWindowToLocal = graph-space
func MapWindowToLocal(x, y float64, offsets imgui.Vec2) (gx, gy float64) {
	gx = float64(offsets.X) + x
	gy = float64(offsets.Y) + y
	return gx, gy
}

// ScrollVelocity adjusts scrolling speed
func ScrollVelocity(scroll float64) float64 {
	sign := 1.0
	if scroll < 0.0 {
		sign = -1.0
	}
	return sign * math.Exp(sign*scroll)
}

func drawHorizontalLine(environment api.IEnvironment, drawList imgui.DrawList,
	y, min, max float64, color imgui.PackedColor) {
	canvasSize := imgui.ContentRegionAvail()
	canvasPos := imgui.CursorScreenPos()

	// ----------------------------------------------------------------
	// Threshold line
	// ----------------------------------------------------------------
	// The sample value needs to be mapped
	uY := MapSampleToUnit(y, min, max)
	// graph space has +Y downward, but the data is oriented as +Y upward
	// so we flip in unit-space.
	uY = 1.0 - uY

	wY := MapUnitToWindow(uY, 0.0, float64(canvasSize.Y))

	_, lY := MapWindowToLocal(0.0, wY, canvasPos)

	wBx := MapUnitToWindow(0.0, 0.0, float64(canvasSize.X))
	wEx := MapUnitToWindow(1.0, 0.0, float64(canvasSize.X))
	lBx, _ := MapWindowToLocal(wBx, 0.0, canvasPos)
	lEx, _ := MapWindowToLocal(wEx, 0.0, canvasPos)

	p1.X = float32(lBx)
	p1.Y = float32(lY)
	p2.X = float32(lEx)
	p2.Y = float32(lY)
	drawList.AddLine(p1, p2, color)
}
