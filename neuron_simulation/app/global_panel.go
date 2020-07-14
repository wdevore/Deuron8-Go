package main

import (
	"fmt"

	"github.com/inkyblackness/imgui-go/v2"
)

var (
	textBuffer = ""
	duration   = int32(0)
	timeScale  = int32(0)
)

func buildGlobalPanel() {
	imgui.Begin("Global Panel")

	if imgui.CollapsingHeader("GlobalHdr") {
		imgui.PushItemWidth(50)
		entered := imgui.InputTextV(
			"Active Sim", &textBuffer,
			imgui.InputTextFlagsEnterReturnsTrue|
				imgui.InputTextFlagsCharsDecimal|
				imgui.InputTextFlagsCharsNoBlank,
			textCallback)

		if entered {
			fmt.Println("Activated Sim: ", textBuffer)
		}
		imgui.PopItemWidth()

		imgui.PushItemWidth(100)
		entered = imgui.InputIntV("Duration", &duration, 1, 100, imgui.InputTextFlagsEnterReturnsTrue)
		if entered {
			fmt.Println("Duration: ", duration)
		}
		imgui.PopItemWidth()

		imgui.PushItemWidth(100)
		imgui.SameLineV(100, 100)
		entered = imgui.InputIntV("Time Scale", &timeScale, 1, 100,
			imgui.InputTextFlagsEnterReturnsTrue|imgui.InputTextFlagsCharsDecimal)
		if entered {
			fmt.Println("Time Scale: ", timeScale)
		}
	}

	imgui.End()

}

func textCallback(data imgui.InputTextCallbackData) int32 {
	fmt.Println(data)
	return 0
}
