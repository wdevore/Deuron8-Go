package gui

import (
	"fmt"
	"strconv"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

// BuildPoissonPanel is an embedded panel inside the parent Simulation panel
func BuildPoissonPanel(environment api.IEnvironment) {
	if imgui.CollapsingHeader("Poisson") {
		sim := environment.Sim()

		simData, _ := sim.Data().(*model.SimJSON)

		imgui.PushItemWidth(80)

		textBuffer = fmt.Sprintf("%3.3f", simData.FiringRate)
		entered := imgui.InputTextV(
			"Firing Rate", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Firing Rate: ", fv)
				sim.Changed()
				simData.FiringRate = fv
			}
		}

		textBuffer = fmt.Sprintf("%3.3f", simData.PoissonPatternMin)
		entered = imgui.InputTextV(
			"Pattern Min", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Pattern Min: ", fv)
				sim.Changed()
				simData.PoissonPatternMin = fv
			}
		}

		imgui.SameLineV(200, 10)

		textBuffer = fmt.Sprintf("%3.3f", simData.PoissonPatternMax)
		entered = imgui.InputTextV(
			"Pattern Max", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Pattern Max: ", fv)
				sim.Changed()
				simData.PoissonPatternMax = fv
			}
		}

		imgui.PopItemWidth()
	}
}

// func textCallback(data imgui.InputTextCallbackData) int32 {
// 	fmt.Println(data)
// 	return 0
// }
