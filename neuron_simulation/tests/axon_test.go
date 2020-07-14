package axontests

// go test -count=1 -v axon_test.go

import (
	"fmt"
	"testing"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/cell"
)

func TestRuns(t *testing.T) {

	axon := cell.NewAxonShiftDelay(0)
	shfAxon := axon.(*cell.AxonShiftDelay)

	if axon.Output() == 1 {
		t.Error("Incorrect output")
	}

	saxon := fmt.Sprintf("%s", shfAxon.Register())

	if saxon != "0000000000000000000000000000000000000000000000000000000000000000" {
		t.Error("Incorrect value")
	}

	axon.Input(1)

	if axon.Output() == 0 {
		t.Error("Incorrect output")
	}

	saxon = fmt.Sprintf("%s", shfAxon.Register())
	if saxon != "0000000000000000000000000000000000000000000000000000000000000001" {
		t.Error("Incorrect value")
	}

	axon.Step()

	saxon = fmt.Sprintf("%s", shfAxon.Register())
	if saxon != "0000000000000000000000000000000000000000000000000000000000000001" {
		t.Error("Incorrect value")
	}

	axon = cell.NewAxonShiftDelay(1 << 1)
	axon.Input(1)
	if axon.Output() == 1 {
		t.Error("Incorrect output")
	}
	fmt.Println(axon)

	axon.Step()
	if axon.Output() == 0 {
		t.Error("Incorrect output")
	}
	fmt.Println(axon)

	axon.Step()
	if axon.Output() == 1 {
		t.Error("Incorrect output")
	}
	fmt.Println(axon)

	fmt.Println("Testing delay 8")
	axon = cell.NewAxonShiftDelay(1 << 8)
	axon.Input(1)

	// We don't expect an output until the 8th step
	for i := 0; i < 7; i++ {
		axon.Step()
		if axon.Output() == 1 {
			t.Error(fmt.Sprintf("Incorrect value at %d", i))
		}
	}

	// Here we expect a "1"
	axon.Step()
	if axon.Output() == 0 {
		t.Error("Incorrect output")
	}
	fmt.Println(axon)

	axon.Input(1)
	fmt.Println(axon)
	axon.Step()
	fmt.Println(axon)
	axon.Step()
	fmt.Println(axon)
	axon.Step()
	fmt.Println(axon)

}
