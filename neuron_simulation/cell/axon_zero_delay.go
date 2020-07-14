package cell

import (
	"fmt"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

// AxonZeroDelay takes an input from the soma and routes it to one or more
// synapses on one or more neurons.
type AxonZeroDelay struct {
	Axon
	input  int
	output int
}

// NewAxonZeroDelay creates an AxonZeroDelay.
func NewAxonZeroDelay() api.IAxon {
	o := new(AxonZeroDelay)
	return o
}

// Reset AxonZeroDelay
func (a *AxonZeroDelay) Reset() {
	a.input = 0
	a.output = 0
}

// Set forces both input and output to a value
func (a *AxonZeroDelay) Set(v int) {
	a.input = v
	a.output = v
}

// Output of AxonZeroDelay that terminates with synapses
func (a *AxonZeroDelay) Output() int {
	return a.output
}

// Input from Soma to AxonZeroDelay's hillock.
func (a *AxonZeroDelay) Input(spike int) {
	a.input = spike
}

// Step output
func (a *AxonZeroDelay) Step() {
	a.output = a.input
}

func (a AxonZeroDelay) String() string {
	return fmt.Sprintf("Input(%d) -> Output(%d)", a.input, a.output)
}
