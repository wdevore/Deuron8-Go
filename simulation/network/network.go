package network

import (
	"github.com/wdevore/Deuron8-Go/api"
	logg "github.com/wdevore/Deuron8-Go/log"
)

type network struct {
	// 2D Grid of neurons
	neurons [][]api.INeuron
}

// New constructs an IVisStimInput object
func NewNetwork() api.INetwork {
	o := new(network)
	return o
}

func (n *network) Load(config api.IConfig) {
	// If we are resuming then we need to reconstruct the environment prior.
	// Otherwise we prepare the environment and then start the simulation.

	// Are we resuming from a previous simulation that was paused?
	if config.ExitState() == "Paused" {
		logg.API.LogInfo("Network loading previous state...")
	} else {
		logg.API.LogInfo("Constructing new Network...")
	}
}

func (n *network) Save() {
}
