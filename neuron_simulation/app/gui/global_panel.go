package gui

import (
	"fmt"
	"strconv"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/graphs"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

var (
	duration  = int32(0)
	timeScale = int32(0)
	maxValue  = float64(0)
	minValue  = float64(0)
)

// BuildGlobalPanel ...
func BuildGlobalPanel(environment api.IEnvironment) {
	imgui.SetNextWindowPos(imgui.Vec2{X: 715, Y: 20.0})

	imgui.Begin("Global Panel")
	config := environment.Config()

	moData, _ := config.Data().(*model.ConfigJSON)

	imgui.PushItemWidth(100)
	duration = int32(moData.Duration)
	entered := imgui.InputIntV("Duration", &duration, 1, 100, imgui.InputTextFlagsEnterReturnsTrue)
	if entered {
		config.Changed()
		fmt.Println("Duration: ", duration)
		moData.Duration = int(duration)
	}
	imgui.PopItemWidth()
	// --------------------------------------------------------------------

	imgui.PushItemWidth(100)
	imgui.SameLineV(100, 100)
	timeScale = int32(moData.TimeScale)
	entered = imgui.InputIntV("Time Scale", &timeScale, 1, 100,
		imgui.InputTextFlagsEnterReturnsTrue|imgui.InputTextFlagsCharsDecimal)
	if entered {
		config.Changed()
		fmt.Println("Time Scale: ", timeScale)
		moData.TimeScale = int(timeScale)
	}

	// --------------------------------------------------------------------
	rangeStart := int32(moData.RangeStart)
	rangeEnd := int32(moData.RangeEnd)
	duration := int32(moData.Duration)

	changedS := imgui.DragIntV("RangeStart##1", &rangeStart, 1.0, 0, int32(moData.RangeEnd), "%d")

	imgui.SameLine()

	changedE := imgui.DragIntV("RangeEnd##1", &rangeEnd, 1.0, rangeStart, duration, "%d")

	if changedS || changedE {
		if rangeStart < rangeEnd {
			config.Changed()
			moData.RangeStart = int(rangeStart)
			moData.RangeEnd = int(rangeEnd)
		}
	}

	scrollVelocity := float32(moData.Scroll)

	changed := imgui.SliderFloatV("Scroll Velocity", &scrollVelocity, -5.0, 5.0, "%.2f", 1.0)
	if changed {
		moData.Scroll = float64(scrollVelocity)
	}

	velocity := graphs.ScrollVelocity(moData.Scroll)
	rangeDx := rangeEnd - rangeStart

	if moData.Scroll < 0 {
		rangeStart += int32(velocity)
		// Left
		if rangeStart > 0 {
			rangeEnd = rangeStart + rangeDx
		} else {
			rangeStart = 0
			rangeEnd = rangeStart + rangeDx
		}
		config.Changed()
		moData.RangeStart = int(rangeStart)
		moData.RangeEnd = int(rangeEnd)
	} else if moData.Scroll > 0 {
		rangeEnd += int32(velocity)
		if rangeEnd < duration {
			rangeStart = rangeEnd - rangeDx
		} else {
			rangeEnd = duration
			rangeStart = rangeEnd - rangeDx
		}
		config.Changed()
		moData.RangeStart = int(rangeStart)
		moData.RangeEnd = int(rangeEnd)
	}

	// If above slider is released we clear the velocity.
	if !imgui.IsItemActive() {
		moData.Scroll = 0.0
	}

	sim := environment.Sim()
	simData, _ := sim.Data().(*model.SimJSON)

	imgui.PushItemWidth(300)
	slide := imgui.SliderInt("Synapse", &activeSynSlide, 0, int32(simData.Synapses)-1)
	if slide {
		sim.Changed()
		simData.ActiveSynapse = int(activeSynSlide)
	}
	imgui.PopItemWidth()

	// --------------------------------------------------------------------
	textBuffer = fmt.Sprintf("%0.2f", minValue)
	entered = imgui.InputTextV(
		"Min Value", &textBuffer,
		imgui.InputTextFlagsEnterReturnsTrue|
			imgui.InputTextFlagsCharsDecimal|
			imgui.InputTextFlagsCharsNoBlank,
		nil)

	if entered {
		fv, err := strconv.ParseFloat(textBuffer, 64)
		if err == nil {
			minValue = fv
		}
	}

	imgui.SameLine()

	textBuffer = fmt.Sprintf("%0.2f", maxValue)
	entered = imgui.InputTextV(
		"Max Value", &textBuffer,
		imgui.InputTextFlagsEnterReturnsTrue|
			imgui.InputTextFlagsCharsDecimal|
			imgui.InputTextFlagsCharsNoBlank,
		nil)

	if entered {
		fv, err := strconv.ParseFloat(textBuffer, 64)
		if err == nil {
			maxValue = fv
		}
	}

	opened := imgui.BeginCombo("Randomizer", "Synapse")
	if opened {
		items := []string{
			"Weights", "TaoP", "TaoN", "TaoI", "Mu",
			"Distance", "Lambda", "Ama", "Amb",
			"Alpha", "LearningRateSlow", "LearningRateFast",
		}
		currentItem := int32(len(items) + 1) // default to no item selected
		changed = imgui.ListBox("", &currentItem, items)

		if changed {
			switch currentItem {
			case 0: // Weights
				environment.SetParms(fmt.Sprintf("Weight,%0.3f,%0.3f", minValue, maxValue))
				environment.IssueCmd("randomizer")
			}
		}

		imgui.EndCombo()
	}

	imgui.End()

}
