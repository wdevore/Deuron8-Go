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
)

// BuildSynapsePanel is an embedded panel inside the parent Simulation panel
func BuildSynapsePanel(environment api.IEnvironment) {
	if imgui.CollapsingHeader("Synapse") {
		sim := environment.Sim()
		synIMod := environment.SynapticModel()
		synMod := synIMod.Data().(*model.SynapsesJSON)

		simData, _ := sim.Data().(*model.SimJSON)
		activeSyn = simData.ActiveSynapse
		synapses := synMod.Synapses
		activeSynapse := synapses[activeSyn]

		imgui.PushItemWidth(80)

		// Row 1 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", activeSynapse.Alpha)
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
				synIMod.Changed()
				if applyToAll {
					for _, syn := range synapses {
						syn.Alpha = fv
					}
				} else {
					activeSynapse.Alpha = fv
				}
			}
		}
		imgui.SameLineV(150, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", activeSynapse.Ama)
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
				synIMod.Changed()
				if applyToAll {
					for _, syn := range synapses {
						syn.Ama = fv
					}
				} else {
					activeSynapse.Ama = fv
				}
			}
		}
		imgui.SameLineV(300, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", activeSynapse.Amb)
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
				synIMod.Changed()
				if applyToAll {
					for _, syn := range synapses {
						syn.Amb = fv
					}
				} else {
					activeSynapse.Amb = fv
				}
			}
		}

		// Row 2 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", activeSynapse.Lambda)
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
				synIMod.Changed()
				if applyToAll {
					for _, syn := range synapses {
						syn.Lambda = fv
					}
				} else {
					activeSynapse.Lambda = fv
				}
			}
		}

		imgui.SameLineV(150, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", activeSynapse.LearningRateFast)
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
				synIMod.Changed()
				if applyToAll {
					for _, syn := range synapses {
						syn.LearningRateFast = fv
					}
				} else {
					activeSynapse.LearningRateFast = fv
				}
			}
		}

		imgui.SameLineV(350, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", activeSynapse.LearningRateSlow)
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
				synIMod.Changed()
				if applyToAll {
					for _, syn := range synapses {
						syn.LearningRateSlow = fv
					}
				} else {
					activeSynapse.LearningRateSlow = fv
				}
			}
		}

		// Row 3 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", activeSynapse.Mu)
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
				synIMod.Changed()
				if applyToAll {
					for _, syn := range synapses {
						syn.Mu = fv
					}
				} else {
					activeSynapse.Mu = fv
				}
			}
		}

		imgui.SameLineV(150, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", activeSynapse.TaoI)
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
				synIMod.Changed()
				if applyToAll {
					for _, syn := range synapses {
						syn.TaoI = fv
					}
				} else {
					activeSynapse.TaoI = fv
				}
			}
		}

		imgui.SameLineV(300, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", activeSynapse.TaoN)
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
				synIMod.Changed()
				if applyToAll {
					for _, syn := range synapses {
						syn.TaoN = fv
					}
				} else {
					activeSynapse.TaoN = fv
				}
			}
		}

		// Row 4 *****************************************************
		textBuffer = fmt.Sprintf("%3.3f", activeSynapse.TaoP)
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
				synIMod.Changed()
				if applyToAll {
					for _, syn := range synapses {
						syn.TaoP = fv
					}
				} else {
					activeSynapse.TaoP = fv
				}
			}
		}

		imgui.SameLineV(150, 10)
		// ----------------------------------------------------------
		textBuffer = fmt.Sprintf("%3.3f", activeSynapse.W)
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
				synIMod.Changed()
				if applyToAll {
					for _, syn := range synapses {
						syn.W = fv
					}
				} else {
					activeSynapse.W = fv
				}
			}
		}

		imgui.PopItemWidth()
	}
}
