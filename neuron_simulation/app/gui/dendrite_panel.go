package gui

import (
	"fmt"
	"strconv"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

// BuildNeuronPanel is an embedded panel inside the parent Simulation panel
func BuildDendritePanel(config, sim api.IModel) {
	if imgui.CollapsingHeader("Dendrite") {
		simData, _ := sim.Data().(*model.SimJSON)
		dendrite := simData.Neuron.Dendrites

		imgui.PushItemWidth(80)

		// Row 1 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", dendrite.TaoEff)
		entered := imgui.InputTextV(
			"TaoEff", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("TaoEff: ", fv)
				sim.Changed()
				dendrite.TaoEff = fv
			}
		}
		imgui.SameLineV(200, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", dendrite.Length)
		entered = imgui.InputTextV(
			"Length", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Length: ", fv)
				sim.Changed()
				dendrite.Length = fv
			}
		}
		imgui.SameLineV(400, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", dendrite.MinPSPValue)
		entered = imgui.InputTextV(
			"MinPSP", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("MinPSP: ", fv)
				sim.Changed()
				dendrite.MinPSPValue = fv
			}
		}

		imgui.PopItemWidth()
	}
}
