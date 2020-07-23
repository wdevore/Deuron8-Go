package gui

import (
	"fmt"
	"os"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

var (
	autosave bool
)

// BuildMenuBar ...
func BuildMenuBar(environment api.IEnvironment) {
	// ---------------------------------------------------------
	// Build the application GUI
	// ---------------------------------------------------------
	config := environment.Config()
	sim := environment.Sim()

	if imgui.BeginMainMenuBar() {
		if imgui.BeginMenu("System") {
			if imgui.MenuItem("Exit") {
				// Save application property settings.
				config.Save()
				environment.IssueCmd("killSim")
				os.Exit(0)
			}
			if imgui.MenuItem("Kill sim thread") {
				environment.IssueCmd("killSim")
			}
			imgui.EndMenu()
		}
		if imgui.BeginMenu("Simulation") {
			if imgui.MenuItem("Load") {
				fmt.Println("Loading simulation properties")
				sim.Load()
			}
			if imgui.MenuItem("Save") {
				fmt.Println("Saving simulation properties")
				sim.Save()
			}
			imgui.EndMenu()
		}
		if imgui.BeginMenu("Settings") {
			moData, _ := config.Data().(*model.ConfigJSON)

			changed := imgui.Checkbox("Autosave", &moData.AutoSave)
			if changed {
				config.Changed()
				if moData.AutoSave {
					fmt.Println("AutoSave enabled")
				} else {
					fmt.Println("AutoSave disabled")
				}
			}

			changed = imgui.Checkbox("Step Enabled", &moData.StepEnabled)
			if changed {
				config.Changed()
				if moData.StepEnabled {
					fmt.Println("Step enabled")
				} else {
					fmt.Println("Step disabled")
				}
			}
			imgui.EndMenu()
		}

		imgui.EndMainMenuBar()
	}
}
