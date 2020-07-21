package network

import (
	"github.com/wdevore/Deuron8-Go/api"
	logg "github.com/wdevore/Deuron8-Go/log"
)

type neuron struct {
	//
}

// New constructs an INeuron object
func NewNeuron() api.INeuron {
	o := new(neuron)
	return o
}

func (n *neuron) Integrate() float64 {
	return 0.0
}

func (n *neuron) Next() {
}

func (n *neuron) Load(config api.IConfig) {
	// If we are resuming then we need to reconstruct the environment prior.
	// Otherwise we prepare the environment and then start the simulation.

	// Are we resuming from a previous simulation that was paused?
	if config.ExitState() == "Paused" {
	} else {
		logg.API.LogInfo("Constructing new Network...")
	}
}

func (n *neuron) Save() {
}
