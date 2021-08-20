package tests

// To Run me:
// go test -count=1 -v model_test.go

import (
	"testing"
)

func TestMain(t *testing.T) {
	simTest(t)
}

func modelTest(t *testing.T) {

	// cmodel := model.NewModel("../../", "jsondata/config.json")

	// mdata, ok := cmodel.Data().(*model.ConfigJSON)

	// if ok {
	// 	fmt.Println(mdata.Simulation)
	// }
}

func simTest(t *testing.T) {

	// cmodel := model.NewSimModel("../../", "jsondata/sim_1.json")

	// mdata, ok := cmodel.Data().(*model.SimJSON)

	// if ok {
	// 	fmt.Println(mdata.Neuron.Dendrites.Compartments[0].Synapses[0].TaoP)
	// }
}
