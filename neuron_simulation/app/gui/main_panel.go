package gui

import (
	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

// BuildMainPanel ...
func BuildMainPanel(environment api.IEnvironment) {
	// imgui.SetNextWindowPos(imgui.Vec2{X: 0.0, Y: 20.0})
	moData, _ := environment.Config().Data().(*model.ConfigJSON)

	imgui.Begin("Main Panel")

	if imgui.Button("Simulate") {
		// Transfer any changed simulation-properties to simulation

		// Reset samples for this new simulation pass

		// Now run simulation
		if moData.StepEnabled {
			environment.IssueCmd("step")
		} else {
			environment.IssueCmd("run")
		}
	}

	imgui.SameLine()
	if imgui.Button("Reset Simulator") {
		environment.IssueCmd("reset")
	}

	imgui.SameLine()
	if imgui.Button("Stop") {
		environment.IssueCmd("stop")
	}

	imgui.End()
}
