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

	environment api.IEnvironment

	simJ     *model.SimJSON
	synJ     *model.SynapsesJSON
	simModel api.IModel
	samples  api.ISamples

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
	// distance float64
}

// NewSynapse creates a new synapse
func NewSynapse(environment api.IEnvironment,
	soma api.ISoma, dendrite api.IDendrite, compartment api.ICompartment,
	id int) api.ISynapse {
	o := new(Synapse)
	o.environment = environment

	simModel := environment.Sim()
	o.simModel = simModel
	o.samples = environment.Samples()
	o.soma = soma
	o.dendrite = dendrite
	o.compartment = compartment
	o.excititory = true // default to excite type
	o.preT = initialPreT
	o.id = id

	simJ, ok := simModel.Data().(*model.SimJSON)

	if !ok {
		panic("Synapse: can't cast simModel to model.SimJSON")
	}

	o.simJ = simJ
	o.synJ, ok = environment.Synapses().Data().(*model.SynapsesJSON)

	if !ok {
		panic("Synapse: can't cast simModel to model.SynapsesJSON")
	}

	compartment.AddSynapse(o)

	return o
}

// Initialize pre configures synapse
func (s *Synapse) Initialize(useDefaults bool) {
	s.Reset()

	// Calc this synapses's reaction to the AP based on its
	// distance from the soma.
	if useDefaults {
		defs := s.simJ.Neuron.Dendrites.Compartments[0].SynapseDefaults

		// Copy defaults to a new model-synapse
		synJ := &model.SynapseJSON{
			TaoP:             defs.TaoP,
			TaoN:             defs.TaoN,
			Mu:               defs.Mu,
			Distance:         defs.Distance,
			ID:               s.id,
			Lambda:           defs.Lambda,
			Amb:              defs.Amb,
			W:                defs.W,
			Alpha:            defs.Alpha,
			LearningRateSlow: defs.LearningRateSlow,
			LearningRateFast: defs.LearningRateFast,
			TaoI:             defs.TaoI,
			Ama:              defs.Ama,
		}
		s.synJ.Synapses = append(s.synJ.Synapses, synJ)

		s.distanceEfficacy = s.dendrite.APEfficacy(defs.Distance)
	} else {
		syn := s.synJ.Synapses[s.id]
		s.distanceEfficacy = s.dendrite.APEfficacy(syn.Distance)
	}
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
	s.wMin = comp.WeightMin
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
	syn := s.synJ.Synapses[s.id]

	// Calc psp based on current dynamics: (t - preT). As dt increases
	// psp will decrease asymtotically to zero.
	dt := float64(t) - s.preT

	dwD := 0.0
	dwP := 0.0
	updateWeight := false

	// The output of the stream is the input to this synapse.
	if s.stream.Output() == 1 {
		// A spike has arrived on the input to this synapse.
		// fmt.Printf("(%d) at %d\n", s.id, t)

		if s.excititory {
			s.surge = s.psp + syn.Ama*math.Exp(-s.psp/syn.TaoP)
		} else {
			s.surge = s.psp + syn.Ama*math.Exp(-s.psp/syn.TaoN)
		}

		// #######################################
		// Depression LTD
		// #######################################
		// Read post trace and adjust weight accordingly.
		dwD = s.prevEffTrace * s.weightFactor(false, syn) * s.soma.APFast()

		s.prevEffTrace = s.efficacy(dt, syn)

		s.preT = float64(t)
		dt = 0.0

		updateWeight = true
	}

	if s.excititory {
		s.psp = s.surge * math.Exp(-dt/syn.TaoP)
	} else {
		s.psp = s.surge * math.Exp(-dt/syn.TaoN)
	}

	// fmt.Printf("dt(%0.3f)|t(%d) surge:%0.3f, exp:%0.3f, psp:%0.3f|\n", dt, t, s.surge, math.Exp(-dt/syn.TaoP), s.psp)

	// If an AP occurred (from the soma) we read the current psp value and add it to the "w"
	if s.soma.Output() == 1 {
		// #######################################
		// Potentiation LTP
		// #######################################
		// Read pre trace (aka psp) and slow AP trace for adjusting weight accordingly.
		//     Post efficacy                                          weight dependence                 triplet sum
		wf := s.weightFactor(true, syn)
		// fmt.Printf("wf:%0.3f", wf)
		dwP = s.soma.EfficacyTrace() * s.distanceEfficacy * wf * (s.psp + s.soma.ApSlowPrior())
		updateWeight = true
	}

	// Finally update the weight.
	if updateWeight {
		s.w = math.Max(math.Min(s.w+dwP-dwD, s.wMax), s.wMin)
	}

	// Return the "value" of this synapse for this "t"
	if s.excititory {
		value = s.psp * s.w * syn.Distance
	} else {
		value = -s.psp * s.w * syn.Distance // is inhibitory
	}

	// Collect this synapse' values at this time step
	s.samples.CollectSynapse(s, s.id, t)

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

// SetStream attaches a spike stream.
func (s *Synapse) SetStream(stream api.IBitStream) {
	s.stream = stream
}

// =============================================================
// Sampling access
// =============================================================

// ID ...
func (s *Synapse) ID() int {
	return s.id
}

// SetID ...
func (s *Synapse) SetID(id int) {
	s.id = id
}

// Weight ...
func (s *Synapse) Weight() float64 {
	return s.w
}

// Surge ...
func (s *Synapse) Surge() float64 {
	return s.surge
}

// Psp ...
func (s *Synapse) Psp() float64 {
	return s.psp
}

// Input ...
func (s *Synapse) Input() int {
	return s.stream.Output()
}
