package cell

import (
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

// Axon takes an input from the soma and routes it to one or more
// synapses on one or more neurons.
// a) Just poisson input
// b) Poisson and Stimulus
type Axon struct {
	length float32
	input  int
	output int
}

// NewAxon creates an axon.
func NewAxon() api.IAxon {
	o := new(Axon)

	return o
}

// Reset axon
func (a *Axon) Reset() {
	a.input = 0
	a.output = 0
}

// Output of axon that terminates with synapses
func (a *Axon) Output() int {
	return a.output
}

// Input from Soma to axon's hillock.
func (a *Axon) Input(spike int) {
	a.input = spike
}

// Step output
func (a *Axon) Step() {
	// Instantanious traversal down axon.
	// Other axon types would implement a delay or shift register.
	a.output = a.input
}
