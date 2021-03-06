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
	imgui.SetNextWindowPos(imgui.Vec2{X: 1390, Y: 20.0})

	imgui.Begin("Simulation")
	sim := environment.Sim()

	if imgui.CollapsingHeader("Simulation Global") {
		simData, _ := sim.Data().(*model.SimJSON)
		moData, _ := environment.Config().Data().(*model.ConfigJSON)

		imgui.PushItemWidth(80)
		textBuffer = fmt.Sprintf("%d", moData.StimulusScaler)
		entered := imgui.InputTextV(
			"Stim Scaler", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseInt(textBuffer, 10, 64)
			if err == nil {
				fmt.Println("Stimulus Scaler: ", fv)
				environment.Config().Changed()
				moData.StimulusScaler = int(fv)
				// Update all Stimulus streams in the simulator
				environment.SetParms("StimulusScaler")
				environment.IssueCmd("propertyChange")
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

		imgui.PopItemWidth()
	}

	BuildPoissonPanel(environment)

	BuildNeuronPanel(environment)

	BuildDendritePanel(environment)

	BuildCompartmentPanel(environment)

	BuildSynapsePanel(environment)

	imgui.End()
}

// func textCallback(data imgui.InputTextCallbackData) int32 {
// 	fmt.Println(data)
// 	return 0
// }
