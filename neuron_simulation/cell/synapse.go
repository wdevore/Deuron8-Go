package cell

import "github.com/wdevore/Deuron8-Go/neuron_simulation/api"

const initialPreT = 0.0 // -1000000000.0

// Synapse is part of a compartment and dendrite
type Synapse struct {
	soma        api.ISoma
	dendrite    api.IDendrite
	compartment api.ICompartment

	id int64

	// true = excititory, false = inhibitory
	excititory bool

	// The weight of the synapse
	w float64

	wMax float64
	wMin float64

	// The stream (aka Merger) that feeds into this synapse
	stream api.IBitStream

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// Surge
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// Surge base value
	amb float64

	// Surge peak
	ama float64

	// Surge window
	tsw float64

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// new surge ion concentration
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// concentration base. We should always have a minimum concentration
	// as a result of a spike
	// Surge is calculated at the arrival of a spike
	// surge = amb - ama*e^(-psp/tsw) == rising curve
	surge float64

	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

	// The time-mark at which a spike arrived at a synapse
	preT float64

	// The current ion concentration
	psp float64

	// =============================================================
	// Learning rules:
	// =============================================================
	// Depression pair-STDP, Potentiation is triplet.
	// "tao"s control the rate of decay. Larger values means a slower decay.
	// Smaller values equals a sharper decay.
	// -----------------------------------

	// denominator, positive window time decay
	taoP float64

	// denominator, negative window time decay
	taoN float64

	// Ratio of mRate/taoX
	tao float64

	// -----------------------------------
	// Weight dependence
	// -----------------------------------
	// F-(w) = λ⍺w^µ, F+(w) = λ(1-w)^µ
	mu     float64 // µ
	lambda float64 // λ
	alpha  float64 // ⍺

	// -----------------------------------
	// Suppression
	// -----------------------------------
	taoI         float64
	prevEffTrace float64

	learningRateSlow float64 // Unused
	learningRateFast float64 // Unused

	// -----------------------------------
	// Fall off
	// -----------------------------------
	distanceEfficacy float64

	// The farther the synapse is from the Soma the less of an influence
	// this synapse has. The function can either be linear or non-linear.
	// The default is linear.
	// Note: no matter how far a synapse is it will still have an influence
	// otherwise it's useless.
	// distance's value = 1.0 for synapses closest to soma. The farther out
	// the value reaches a minimum of around ~0.25.
	// distance is multiplied into soma.psp.
	distance float64
}

// NewSynapse creates a new synapse
func NewSynapse(soma api.ISoma, dendrite api.IDendrite, compartment api.ICompartment) api.ISynapse {
	o := new(Synapse)
	o.soma = soma
	o.dendrite = dendrite
	o.compartment = compartment
	o.excititory = true // default to excite type
	o.preT = initialPreT
	o.id = 0

	// add_synapse!(compartment, o)

	return o
}

// Initialize pre configures synapse
func (s *Synapse) Initialize() {
	// Focus the model on the correct synapse.
}

// Reset resets for another sim pass
func (s *Synapse) Reset() {
	s.prevEffTrace = 1.0
	s.surge = 0.0
	s.psp = 0.0
	s.preT = 0.0

	// Reset weights back to best guess values.
	s.wMax = s.compartment.weight_max
	s.w = s.wMax / s.compartment.weight_divisor

}

// SetType Inhibit=false, excititory=true
func (s *Synapse) SetType(sType bool) {
	s.excititory = sType
}

// PreIntegrate is called prior integration
func (s *Synapse) PreIntegrate() {}

// Integrate is the actual integration
func (s *Synapse) Integrate() {}

// PostIntegrate is called after integration
func (s *Synapse) PostIntegrate() {}

// TripleIntegration advanced
// =============================================================
// Triplet:
// =============================================================
// Pre trace, Post slow and fast traces.
//
// Depression: fast post trace with at pre spike
// Potentiation: slow post trace at post spike
func (s *Synapse) TripleIntegration(spanT, t int64) {}

// Efficacy : each spike of pre-synaptic neuron j sets the presynaptic spike
// efficacy j to 0
// whereafter it recovers exponentially to 1 with a time constant
// τj = toaJ
// In other words, the efficacy of a spike is suppressed by
// the proximity of a previous spike.
func (s *Synapse) Efficacy(dt float64) {}

// WeightFactor mu = 0.0 = additive, mu = 1.0 = multiplicative
func (s *Synapse) WeightFactor(potentiation bool) {}

// SetStream attaches a spike stream
func (s *Synapse) SetStream(stream api.IBitStream) {
	s.stream = stream
}
