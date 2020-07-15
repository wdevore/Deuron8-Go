package gui

import (
	"fmt"
	"strconv"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

// BuildNeuronPanel is an embedded panel inside the parent Simulation panel
func BuildNeuronPanel(config, sim api.IModel) {
	if imgui.CollapsingHeader("Neuron") {
		simData, _ := sim.Data().(*model.SimJSON)
		neuron := simData.Neuron

		imgui.PushItemWidth(80)

		// Row 1 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", neuron.RefractoryPeriod)
		entered := imgui.InputTextV(
			"RefractoryPeriod", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("RefractoryPeriod: ", fv)
				sim.Changed()
				neuron.RefractoryPeriod = fv
			}
		}
		imgui.SameLineV(200, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", neuron.Threshold)
		entered = imgui.InputTextV(
			"Threshold", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Threshold: ", fv)
				sim.Changed()
				neuron.Threshold = fv
			}
		}
		imgui.SameLineV(400, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", neuron.APMax)
		entered = imgui.InputTextV(
			"APMax", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("APMax: ", fv)
				sim.Changed()
				neuron.APMax = fv
			}
		}

		// Row 2 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", neuron.FastSurge)
		entered = imgui.InputTextV(
			"Fast Surge", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Fast Surge: ", fv)
				sim.Changed()
				neuron.FastSurge = fv
			}
		}
		imgui.SameLineV(200, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", neuron.SlowSurge)
		entered = imgui.InputTextV(
			"Slow Surge", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Slow Surge: ", fv)
				sim.Changed()
				neuron.SlowSurge = fv
			}
		}

		// Row 3 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", neuron.Tao)
		entered = imgui.InputTextV(
			"Tao", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Tao: ", fv)
				sim.Changed()
				neuron.Tao = fv
			}
		}
		imgui.SameLineV(200, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", neuron.TaoJ)
		entered = imgui.InputTextV(
			"Tao J", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Tao J: ", fv)
				sim.Changed()
				neuron.TaoJ = fv
			}
		}
		imgui.SameLineV(400, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", neuron.TaoS)
		entered = imgui.InputTextV(
			"Tao S", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Tao S: ", fv)
				sim.Changed()
				neuron.TaoS = fv
			}
		}

		imgui.PopItemWidth()
	}
}
