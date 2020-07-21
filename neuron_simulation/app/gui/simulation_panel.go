package gui

import (
	"fmt"
	"strconv"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

var (
	actSyn = int32(0)
)

// BuildSimulationPanel ...
func BuildSimulationPanel(environment api.IEnvironment) {
	imgui.Begin("Simulation")
	sim := environment.Sim()

	if imgui.CollapsingHeader("Simulation Global") {
		simData, _ := sim.Data().(*model.SimJSON)

		imgui.PushItemWidth(80)
		textBuffer = fmt.Sprintf("%3.3f", simData.StimulusScaler)
		entered := imgui.InputTextV(
			"Stim Scaler", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Stim Scaler: ", fv)
				sim.Changed()
				simData.StimulusScaler = fv
			}
		}

		imgui.SameLineV(200, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%d", simData.Hertz)
		entered = imgui.InputTextV(
			"Hertz", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			fv, err := strconv.ParseInt(textBuffer, 10, 64)
			if err == nil {
				fmt.Println("Hertz: ", fv)
				sim.Changed()
				simData.Hertz = int(fv)
			}
		}

		imgui.SameLineV(350, 10)
		// ----------------------------------------------------------
		actSyn = int32(simData.ActiveSynapse)
		entered = imgui.InputIntV("Active Synapse", &actSyn, 1, 100,
			imgui.InputTextFlagsEnterReturnsTrue|imgui.InputTextFlagsCharsDecimal)
		if entered {
			if actSyn < 0 {
				actSyn = 0
			} else if actSyn >= int32(simData.Synapses) {
				actSyn = int32(simData.Synapses) - 1
			}
			sim.Changed()
			fmt.Println("Active Synapse: ", actSyn)
			simData.ActiveSynapse = int(actSyn)
		}

		imgui.PopItemWidth()
	}

	BuildPoissonPanel(environment)

	BuildNeuronPanel(environment)

	BuildDendritePanel(environment)

	BuildSynapsePanel(environment)

	imgui.End()
}

// func textCallback(data imgui.InputTextCallbackData) int32 {
// 	fmt.Println(data)
// 	return 0
// }
