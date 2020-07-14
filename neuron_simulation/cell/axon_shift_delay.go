package cell

import (
	"fmt"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

// AxonShiftDelay takes an input from the soma and routes it to one or more
// synapses on one or more neurons.
type AxonShiftDelay struct {
	Axon

	// A single bit is Set at a particular position where the
	// output is sampled.
	//                      /--- output
	//                     |
	// An 8bit example: 00010000
	//                          \-- input
	delay uint64

	// MSB = output ... LSB = input
	register uint64
}

// NewAxonShiftDelay creates an AxonShiftDelay.
func NewAxonShiftDelay(delay uint64) api.IAxon {
	o := new(AxonShiftDelay)
	if delay == 0 {
		delay = 1
	}
	o.delay = delay
	return o
}

// Reset AxonShiftDelay
func (a *AxonShiftDelay) Reset() {
	a.register = 0
}

// SetNoDelay sets a zero delay value
func (a *AxonShiftDelay) SetNoDelay() {
	a.delay = 1
}

// SetToMaxDelay set maximum delay (64 delays)
func (a *AxonShiftDelay) SetToMaxDelay() {
	a.delay = 1 << 63
}

// SetToHalfDelay set half maximum delay (32 delays)
func (a *AxonShiftDelay) SetToHalfDelay() {
	a.delay = 1 << 31
}

// Output of AxonShiftDelay that terminates with synapses
func (a *AxonShiftDelay) Output() int {
	// Capture output first
	if a.delay == 1 {
		return int(a.register)
	}

	if int(a.register&a.delay) > 0 {
		return 1
	}

	return 0
}

// Input from Soma to AxonShiftDelay's hillock.
// The incoming spike enters at position 0 (LSB) and shifts
// towards the MSB
func (a *AxonShiftDelay) Input(spike int) {
	// Place spike in register at LSB
	a.register = a.register | uint64(spike&1)
}

// Step output
func (a *AxonShiftDelay) Step() {
	if a.delay > 1 {
		a.register = a.register << 1
	}
}

// Register returns 64bit string
func (a AxonShiftDelay) Register() string {
	return fmt.Sprintf("%064b", a.register)
}

func (a AxonShiftDelay) String() string {
	return fmt.Sprintf("%064b register\n%064b delay\n%d Output", a.register, a.delay, a.Output())
}
