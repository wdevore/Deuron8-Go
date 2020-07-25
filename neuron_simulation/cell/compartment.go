package cell

import (
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

// Compartment is part of a cell
type Compartment struct {
	soma     api.ISoma
	dendrite api.IDendrite

	simJ     *model.SimJSON
	simModel api.IModel

	// Contains Synapses
	synapses []api.ISynapse
}

// NewCompartment creates a new compartment containing synapses
func NewCompartment(simModel api.IModel, dendrite api.IDendrite, soma api.ISoma) api.ICompartment {
	o := new(Compartment)
	o.synapses = []api.ISynapse{}
	o.soma = soma
	o.dendrite = dendrite

	simJ, ok := simModel.Data().(*model.SimJSON)

	if !ok {
		panic("Compartment: can't cast simModel to SimJSON")
	}

	o.simJ = simJ

	dendrite.AddCompartment(o)

	return o
}

// Initialize compartment
func (c *Compartment) Initialize() {
}

// AddSynapse adds synapse to compartment
func (c *Compartment) AddSynapse(synapse api.ISynapse) {
	c.synapses = append(c.synapses, synapse)
}

// Reset compartment
func (c *Compartment) Reset() {
	for _, synapses := range c.synapses {
		synapses.Reset()
	}
}

// Integrate is the actual integration
func (c *Compartment) Integrate(spanT, t int) (psp, totalWeight float64) {
	psp = 0.0

	for _, synapse := range c.synapses {
		sum, _ := synapse.Integrate(spanT, t)
		// sum, weight := synapse.Integrate(spanT, t)
		psp += sum
		// totalWeight += weight
	}

	return psp, totalWeight
}
