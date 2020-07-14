package cell

import (
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

// Axon takes an input from the soma and routes it to one or more
// synapses on one or more neurons.
// a) Just poisson input
// b) Poisson and Stimulus
type Axon struct {
}

// NewAxon creates an axon.
func NewAxon() api.IAxon {
	o := new(Axon)

	return o
}

// Reset axon
func (a *Axon) Reset() {
}

func (a *Axon) Set(v int) {
}

func (a *Axon) SetNoDelay() {
}

func (a *Axon) SetToMaxDelay() {
}

func (a *Axon) SetToHalfDelay() {
}

// Output of axon that terminates with synapses
func (a *Axon) Output() int {
	return 0
}

// Input from Soma to axon's hillock.
func (a *Axon) Input(spike int) {
}

// Step output
func (a *Axon) Step() {
}
