package gui

import (
	"fmt"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

var (
	duration  = int32(0)
	timeScale = int32(0)
)

// BuildGlobalPanel ...
func BuildGlobalPanel(environment api.IEnvironment) {
	imgui.Begin("Global Panel")
	config := environment.Config()

	// if imgui.CollapsingHeader("GlobalHdr") {
	imgui.PushItemWidth(50)
	entered := imgui.InputTextV(
		"Active Sim", &textBuffer,
		imgui.InputTextFlagsEnterReturnsTrue|
			imgui.InputTextFlagsCharsDecimal|
			imgui.InputTextFlagsCharsNoBlank,
		nil)

	if entered {
		fmt.Println("Activated Sim: ", textBuffer)
	}
	imgui.PopItemWidth()

	moData, _ := config.Data().(*model.ConfigJSON)

	imgui.PushItemWidth(100)
	duration = int32(moData.Duration)
	entered = imgui.InputIntV("Duration", &duration, 1, 100, imgui.InputTextFlagsEnterReturnsTrue)
	if entered {
		config.Changed()
		fmt.Println("Duration: ", duration)
		moData.Duration = int(duration)
	}
	imgui.PopItemWidth()

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
	// }

	imgui.End()

}
