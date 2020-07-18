package gui

import (
	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

// BuildMainPanel ...
func BuildMainPanel(environment api.IEnvironment) {
	// imgui.SetNextWindowPos(imgui.Vec2{X: 0.0, Y: 20.0})

	imgui.Begin("Main Panel")

	if imgui.Button("Simulate") {
		// Transfer any changed simulation-properties to simulation

		// Reset samples for this new simulation pass

		// Now run simulation
		environment.IssueCmd("start")
	}

	imgui.SameLine()
	if imgui.Button("Simulate Once") {
		environment.IssueCmd("once")
	}

	imgui.SameLine()
	if imgui.Button("Stop") {
		environment.IssueCmd("stop")
	}

	imgui.End()
}
