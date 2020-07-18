package gui

import (
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

var (
	textBuffer = ""
)

// BuildGui ...
func BuildGui(environment api.IEnvironment) {

	BuildMenuBar(environment)

	BuildMainPanel(environment)

	BuildGlobalPanel(environment)

	BuildSimulationPanel(environment)
}

// Shutdown saves config
func Shutdown(environment api.IEnvironment) {
	environment.Config().Save()
}
