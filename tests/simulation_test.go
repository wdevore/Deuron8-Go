package tests

import (
	"testing"

	"github.com/wdevore/Deuron8-Go/config"
	logg "github.com/wdevore/Deuron8-Go/log"
	"github.com/wdevore/Deuron8-Go/simulation/network"
)

func init() {
	config.API = config.New("../config/config.json")
	logg.API = logg.New(config.API)
}

func TestSimulation(t *testing.T) {
	nwk := network.NewNetwork()

	nwk.Load(config.API)
}
