package gui

import (
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

var (
	textBuffer = ""
)

// BuildGui ...
func BuildGui(config, sim api.IModel) {

	BuildMenuBar(config, sim)

	BuildMainPanel(config)

	BuildGlobalPanel(config)

	BuildSimulationPanel(config, sim)
}

// Shutdown saves config
func Shutdown(config api.IModel) {
	config.Save()
}
