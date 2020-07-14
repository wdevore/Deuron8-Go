package cell

import (
	"math"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

const initialPreT = 0.0 // -1000000000.0

// Synapse is part of a compartment and dendrite
type Synapse struct {
	soma        api.ISoma
	dendrite    api.IDendrite
	compartment api.ICompartment

	simJ     *model.SimJSON
	simModel api.IModel

	id int

	// true = excititory, false = inhibitory
	excititory bool

	// The weight of the synapse
	w float64

	wMax float64
	wMin float64

	// The stream (aka Merger) that feeds into this synapse
	stream api.IBitStream

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

	// -----------------------------------
	// Weight dependence
	// -----------------------------------
	// F-(w) = λ⍺w^µ, F+(w) = λ(1-w)^µ

	// -----------------------------------
	// Suppression
	// -----------------------------------
	prevEffTrace float64

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
func NewSynapse(simModel api.IModel, soma api.ISoma, dendrite api.IDendrite, compartment api.ICompartment) api.ISynapse {
	o := new(Synapse)
	o.simModel = simModel
	o.soma = soma
	o.dendrite = dendrite
	o.compartment = compartment
	o.excititory = true // default to excite type
	o.preT = initialPreT
	o.id = 0

	simJ, ok := simModel.Data().(*model.SimJSON)

	if !ok {
		panic("Synapse: can't cast simModel to SimJSON")
	}

	o.simJ = simJ

	compartment.AddSynapse(o)

	return o
}

// Initialize pre configures synapse
func (s *Synapse) Initialize() {

	// s.simJ.ActiveSynapse = s.id

	// Calc this synapses's reaction to the AP based on its
	// distance from the soma.
	syn := s.simJ.Neuron.Dendrites.Compartments[0].Synapses[s.id]
	s.distanceEfficacy = s.dendrite.APEfficacy(syn.Distance)
}

// Reset resets for another sim pass
func (s *Synapse) Reset() {
	s.prevEffTrace = 1.0
	s.surge = 0.0
	s.psp = 0.0
	s.preT = 0.0

	// Reset weights back to best guess values.
	comp := s.simJ.Neuron.Dendrites.Compartments[0]
	s.wMax = comp.WeightMax
	s.w = s.wMax / comp.WeightDivisor
}

// SetType Inhibit=false, excititory=true
func (s *Synapse) SetType(sType bool) {
	s.excititory = sType
}

// PreIntegrate is called prior integration
func (s *Synapse) PreIntegrate() {

}

// Integrate is the actual integration
func (s *Synapse) Integrate(spanT, t int) (value, w float64) {
	return s.tripleIntegration(spanT, t)
}

// PostIntegrate is called after integration
func (s *Synapse) PostIntegrate() {
}

// TripleIntegration advanced
// =============================================================
// Triplet:
// =============================================================
// Pre trace, Post slow and fast traces.
//
// Depression: fast post trace with at pre spike
// Potentiation: slow post trace at post spike
func (s *Synapse) tripleIntegration(spanT, t int) (value, w float64) {
	syn := s.simJ.Neuron.Dendrites.Compartments[0].Synapses[s.id]

	// Calc psp based on current dynamics: (t - preT). As dt increases
	// psp will decrease asymtotically to zero.
	dt := float64(t) - s.preT

	dwD := 0.0
	dwP := 0.0
	updateWeight := false

	// The output of the stream is the input to this synapse.
	if s.stream.Output() == 1 {
		// A spike has arrived on the input to this synapse.
		// println("(", t, ") syn: ", syn.id)

		if s.excititory {
			s.surge = s.psp + syn.Ama*math.Exp(-s.psp/syn.TaoP)
		} else {
			s.surge = s.psp + syn.Ama*math.Exp(-s.psp/syn.TaoN)
		}

		// #######################################
		// Depression LTD
		// #######################################
		// Read post trace and adjust weight accordingly.
		dwD = s.prevEffTrace * s.weightFactor(false, &syn) * s.soma.APFast()

		s.prevEffTrace = s.efficacy(dt, &syn)

		s.preT = float64(t)
		dt = 0.0

		updateWeight = true
	}

	if s.excititory {
		s.psp = s.surge * math.Exp(-dt/syn.TaoP)
	} else {
		s.psp = s.surge * math.Exp(-dt/syn.TaoN)
	}

	// If an AP occurred (from the soma) we read the current psp value and add it to the "w"
	if s.soma.Output() == 1 {
		// #######################################
		// Potentiation LTP
		// #######################################
		// Read pre trace (aka psp) and slow AP trace for adjusting weight accordingly.
		//     Post efficacy                                          weight dependence                 triplet sum
		dwP = s.soma.EfficacyTrace() * s.distanceEfficacy * s.weightFactor(true, &syn) * (s.psp + s.soma.ApSlowPrior())
		updateWeight = true
	}

	// Finally update the weight.
	if updateWeight {
		s.w = math.Max(math.Min(s.w+dwP-dwD, s.wMax), s.wMin)
	}

	// Return the "value" of this synapse for this "t"
	if s.excititory {
		value = s.psp * s.w * s.distance
	} else {
		value = -s.psp * s.w * s.distance // is inhibitory
	}

	// Collect this synapse' values at this time step
	s.simModel.Samples().CollectSynapse(s, t)

	return value, s.w
}

// Efficacy : each spike of pre-synaptic neuron j sets the presynaptic spike
// efficacy j to 0
// whereafter it recovers exponentially to 1 with a time constant
// τj = toaJ
// In other words, the efficacy of a spike is suppressed by
// the proximity of a trailing spike.
func (s *Synapse) efficacy(dt float64, syn *model.SynapseJSON) float64 {
	return 1.0 - math.Exp(-dt/syn.TaoI)

}

// WeightFactor mu = 0.0 = additive, mu = 1.0 = multiplicative
func (s *Synapse) weightFactor(potentiation bool, syn *model.SynapseJSON) float64 {
	if potentiation {
		return syn.Lambda * math.Pow(1.0-syn.W/s.wMax, syn.Mu)
	}

	return syn.Lambda * syn.Alpha * math.Pow(syn.W/s.wMax, syn.Mu)
}

// SetStream attaches a spike stream
func (s *Synapse) SetStream(stream api.IBitStream) {
	s.stream = stream
}
