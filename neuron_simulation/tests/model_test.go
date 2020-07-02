package imguitests

// To Run me:
// go test -count=1 -v model_test.go

import (
	"fmt"
	"testing"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

func TestMain(t *testing.T) {
	simTest(t)
}

func modelTest(t *testing.T) {

	cmodel := model.NewModel("../../", "jsondata/config.json")

	mdata, ok := cmodel.Data().(*model.ConfigJSON)

	if ok {
		fmt.Println(mdata.Simulation)
	}
}

func simTest(t *testing.T) {

	cmodel := model.NewSimModel("../../", "jsondata/sim_1.json")

	mdata, ok := cmodel.Data().(*model.SimJSON)

	if ok {
		fmt.Println(mdata.Neuron.Dendrites.Compartments[0].Synapses[0].TaoP)
	}
}
