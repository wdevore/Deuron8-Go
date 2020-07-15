package gui

import (
	"fmt"
	"strconv"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

var (
	activeSyn      = 0
	activeSynSlide = int32(0)
	applyToAll     bool
)

// BuildSynapsePanel is an embedded panel inside the parent Simulation panel
func BuildSynapsePanel(config, sim api.IModel) {
	if imgui.CollapsingHeader("Synapse") {
		simData, _ := sim.Data().(*model.SimJSON)
		activeSyn = simData.ActiveSynapse
		synapse := simData.Neuron.Dendrites.Compartments[0].Synapses[activeSyn]

		imgui.PushItemWidth(80)

		// Row 1 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", synapse.Alpha)
		entered := imgui.InputTextV(
			"Alpha", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Alpha: ", fv)
				sim.Changed()
				synapse.Alpha = fv
			}
		}
		imgui.SameLineV(150, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", synapse.Ama)
		entered = imgui.InputTextV(
			"Ama", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Ama: ", fv)
				sim.Changed()
				synapse.Ama = fv
			}
		}
		imgui.SameLineV(300, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", synapse.Amb)
		entered = imgui.InputTextV(
			"Amb", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Amb: ", fv)
				sim.Changed()
				synapse.Amb = fv
			}
		}

		// Row 2 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", synapse.Lambda)
		entered = imgui.InputTextV(
			"Lambda", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Lambda: ", fv)
				sim.Changed()
				synapse.Lambda = fv
			}
		}

		imgui.SameLineV(150, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", synapse.LearningRateFast)
		entered = imgui.InputTextV(
			"Fast Learn Rate", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Fast Learn Rate: ", fv)
				sim.Changed()
				synapse.LearningRateFast = fv
			}
		}

		imgui.SameLineV(350, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", synapse.LearningRateSlow)
		entered = imgui.InputTextV(
			"Slow Learn Rate", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Slow Learn Rate: ", fv)
				sim.Changed()
				synapse.LearningRateSlow = fv
			}
		}

		// Row 3 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", synapse.Mu)
		entered = imgui.InputTextV(
			"Mu", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Mu: ", fv)
				sim.Changed()
				synapse.Mu = fv
			}
		}

		imgui.SameLineV(150, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", synapse.TaoI)
		entered = imgui.InputTextV(
			"TaoI", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("TaoI: ", fv)
				sim.Changed()
				synapse.TaoI = fv
			}
		}

		imgui.SameLineV(300, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", synapse.TaoN)
		entered = imgui.InputTextV(
			"TaoN", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("TaoN: ", fv)
				sim.Changed()
				synapse.TaoN = fv
			}
		}

		// Row 4 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", synapse.TaoP)
		entered = imgui.InputTextV(
			"TaoP", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("TaoP: ", fv)
				sim.Changed()
				synapse.TaoP = fv
			}
		}

		imgui.SameLineV(150, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", synapse.W)
		entered = imgui.InputTextV(
			"Weight", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			nil)

		if entered {
			sim.Changed()
			fv, err := strconv.ParseFloat(textBuffer, 64)
			if err == nil {
				fmt.Println("Weight: ", fv)
				sim.Changed()
				synapse.W = fv
			}
		}

		// Row 5 *****************************************************
		imgui.PushItemWidth(450)
		slide := imgui.SliderInt("ActiveSynapse", &activeSynSlide, 0, int32(simData.Synapses)-1)
		if slide {
			sim.Changed()
			fmt.Println("slide: ", activeSynSlide)
			simData.ActiveSynapse = int(activeSynSlide)
		}
		imgui.PopItemWidth()

		// TODO complete
		changed := imgui.Checkbox("Apply To All", &applyToAll)
		if changed {
			config.Changed()
			if applyToAll {
				fmt.Println("applyToAll enabled")
			} else {
				fmt.Println("applyToAll disabled")
			}
		}

		imgui.PopItemWidth()
	}
}
