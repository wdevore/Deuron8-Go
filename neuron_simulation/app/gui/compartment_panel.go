package gui

import (
	"fmt"
	"strconv"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

// BuildCompartmentPanel is an embedded panel inside the parent Simulation panel
func BuildCompartmentPanel(environment api.IEnvironment) {
	if imgui.CollapsingHeader("Compartment") {
		sim := environment.Sim()
		simData, _ := sim.Data().(*model.SimJSON)
		compartment := simData.Neuron.Dendrites.Compartments[0]

		imgui.PushItemWidth(80)

		// Row 1 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", compartment.WeightMin)
		entered := imgui.InputTextV(
			"Weight Min", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Weight Max: ", fv)
				sim.Changed()
				compartment.WeightMin = fv
			}
		}

		imgui.SameLineV(200, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", compartment.WeightMax)
		entered = imgui.InputTextV(
			"Weight Max", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Weight Max: ", fv)
				sim.Changed()
				compartment.WeightMax = fv
			}
		}

		imgui.PopItemWidth()
	}
}
