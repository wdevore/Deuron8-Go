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
	running  bool
)

// BuildMenuBar ...
func BuildMenuBar(config api.IModel) {
	// ---------------------------------------------------------
	// Build the application GUI
	// ---------------------------------------------------------
	if imgui.BeginMainMenuBar() {
		if imgui.BeginMenu("System") {
			if imgui.MenuItem("Exit") {
				// Save application property settings.
				config.Save()

				os.Exit(0)
			}
			imgui.EndMenu()
		}
		if imgui.BeginMenu("Simulation") {
			if imgui.MenuItem("Load") {
				fmt.Println("Loading simulation properties")
			}
			if imgui.MenuItem("Save") {
				fmt.Println("Saving simulation properties")
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
			imgui.EndMenu()
		}

		imgui.EndMainMenuBar()
	}
}
