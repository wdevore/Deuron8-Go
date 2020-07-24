package graphs

import (
	"math"

	"github.com/inkyblackness/imgui-go/v2"
)

const (
	SpikePanelHeight = 300
	SurgePanelHeight = 200
	PspPanelHeight   = 200
)

// Lerp returns a the value between min and max given t = 0->1
func Lerp(min, max, t float64) float64 {
	return min*(1.0-t) + max*t
}

// Linear returns 0->1 for a "value" between min and max.
// Generally used to map from view-space to unit-space
func Linear(min, max, value float64) float64 {
	if max < min {
		tmp := max
		max = min
		min = tmp
	}

	if min < 0.0 {
		return 1.0 - (value-max)/(min-max)
	}

	return (value - min) / (max - min)
}

// MapSampleToUnit from sample-space to unit-space where unit-space is 0->1
func MapSampleToUnit(v, min, max float64) float64 {
	return Linear(min, max, v)
}

// MapUnitToWindow from unit-space to window-space
func MapUnitToWindow(v, min, max float64) float64 {
	return Lerp(min, max, v)
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
