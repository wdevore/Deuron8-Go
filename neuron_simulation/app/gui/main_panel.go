package gui

import (
	"fmt"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

// BuildMainPanel ...
func BuildMainPanel(config api.IModel) {
	// imgui.SetNextWindowPos(imgui.Vec2{X: 0.0, Y: 20.0})

	imgui.Begin("Main Panel")

	if imgui.Button("Simulate") {
		// Transfer any changed simulation-properties to simulation

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