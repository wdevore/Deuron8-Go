package main

import (
	"fmt"

	"github.com/inkyblackness/imgui-go/v2"
)

func buildMainPanel() {
	imgui.SetNextWindowPos(imgui.Vec2{0.0, 20.0})

	imgui.Begin("Main Panel")

	if imgui.Button("Simulate") {
		if autosave {
			fmt.Println("Saving model data.")
		}

		// Transfer any changed GUI data to simulation

		// Reset samples for this new simulation pass

		// Now run simulation
		fmt.Println("Running simulation...")
		running = true
	}

	imgui.SameLine()
	if imgui.Button("Stop") {
		fmt.Println("Stopping simulation...")
		running = false
	}

	imgui.End()
}
