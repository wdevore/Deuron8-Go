package simulation

import (
	"github.com/wdevore/Deuron8-Go/api"
	"github.com/wdevore/Deuron8-Go/config"
	logg "github.com/wdevore/Deuron8-Go/log"
	"github.com/wdevore/Deuron8-Go/simulation/network"
)

var debug = 0
var netw api.INetwork

func construct() {
	netw := network.New()

	netw.Load(config.API)
}

// A simulation runs in chunks. This gives the "application" a chance
// to check on user input. A chunk size depends on how long a group
// of simulation steps take.
func simulate() bool {

	debug++
	if debug > 10 {
		return false
	}
	logg.API.LogInfo("Simulating...")

	return true
}
