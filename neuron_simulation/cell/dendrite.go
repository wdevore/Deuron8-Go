package cell

import (
	"math"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

// Dendrite is part of a compartment
type Dendrite struct {
	soma api.ISoma

	simJ     *model.SimJSON
	simModel api.IModel

	// These fields are not used yet.
	taoEff float64

	// Minimum value. Typically 0.0
	minPsp float64

	// Contains Compartments
	compartments []api.ICompartment

	// Average weight over time
	average  float64
	synapses int
}

// NewDendrite creates a new dendrite
func NewDendrite(simModel api.IModel, soma api.ISoma) api.IDendrite {
	o := new(Dendrite)
	o.simModel = simModel

	o.compartments = []api.ICompartment{}
	o.soma = soma

	simJ, ok := simModel.Data().(*model.SimJSON)

	if !ok {
		panic("Dendrite: can't cast simModel to SimJSON")
	}

	o.simJ = simJ

	return o
}

// Initialize dendrite
func (d *Dendrite) Initialize() {
	for _, compartment := range d.compartments {
		compartment.Initialize()
	}
}

// AddCompartment adds compartment
func (d *Dendrite) AddCompartment(compartment api.ICompartment) {
	d.compartments = append(d.compartments, compartment)
}

// Reset dendrite
func (d *Dendrite) Reset() {
	for _, compartment := range d.compartments {
		compartment.Reset()
	}
}

// APEfficacy Calc this synapses's reaction to the AP based on its
// distance from the soma.
func (d *Dendrite) APEfficacy(distance float64) float64 {
	dendrite := d.simJ.Neuron.Dendrites

	if distance < dendrite.Length {
		return 1.0
	}

	return math.Exp(-(dendrite.Length - distance) / dendrite.TaoEff)
}

// Integrate is the actual integration
func (d *Dendrite) Integrate(spanT, t int) (psp float64) {
	totalWeight := 0.0

	for _, compartment := range d.compartments {
		sum, total := compartment.Integrate(spanT, t)
		psp += sum
		totalWeight += total
	}

	psp = math.Max(psp, d.minPsp)

	d.average = totalWeight / float64(d.synapses)

	// Collect this Dendrite' values at this time step
	d.simModel.Samples().CollectDendrite(d, spanT)

	return psp
}
