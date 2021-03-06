package cell

import (
	"math"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

// Soma is the body of a cell
type Soma struct {
	simJ     *model.SimJSON
	simModel api.IModel
	samples  api.ISamples

	// Axon is the output
	axon api.IAxon

	dendrite api.IDendrite

	// Soma threshold. When exceeded an AP is generated.
	threshold float64
	// Post synaptic potential
	psp float64

	// --------------------------------------------------------
	// Action potential
	// --------------------------------------------------------
	// AP can travel back down the dendrite. The value decays
	// with distance.
	apFast      float64 // Fast trace
	apSlow      float64 // Slow trace
	apSlowPrior float64 // Slow trace (t-1)

	// The time-mark of the current AP.
	APt float64
	// The previous time-mark
	preAPt float64
	APMax  float64

	// --------------------------------------------------------
	// STDP
	// --------------------------------------------------------
	// -----------------------------------
	// AP decay
	// -----------------------------------
	// ntao  float64 // fast trace
	// ntaoS float64 // slow trace

	// Fast Surge
	nFastSurge        float64
	nDynFastSurge     float64
	nInitialFastSurge float64

	// Slow Surge
	nSlowSurge        float64
	nDynSlowSurge     float64
	nInitialSlowSurge float64

	// The time-mark at which a spike arrived at a synapse
	preT float64

	refractoryCnt   float64
	refractoryState bool

	// -----------------------------------
	// Suppression
	// -----------------------------------
	ntaoJ         float64
	efficacyTrace float64
}

// NewSoma creates an Soma.
func NewSoma(simModel api.IModel, samples api.ISamples) api.ISoma {
	o := new(Soma)
	o.samples = samples

	simJ, ok := simModel.Data().(*model.SimJSON)

	if !ok {
		panic("Synapse: can't cast simModel to SimJSON")
	}

	o.simJ = simJ

	return o
}

// Initialize soma
func (s *Soma) Initialize() {
	s.dendrite.Initialize()
}

// SetDendrite attaches dendrite
func (s *Soma) SetDendrite(dendrite api.IDendrite) {
	s.dendrite = dendrite
}

// SetAxon attaches axon
func (s *Soma) SetAxon(axon api.IAxon) {
	s.axon = axon
}

// Reset soma
func (s *Soma) Reset() {
	// Initial values to boot the cell
	neuron := s.simJ.Neuron
	s.nSlowSurge = neuron.SlowSurge
	s.nFastSurge = neuron.FastSurge
	s.threshold = neuron.Threshold
	s.APMax = neuron.APMax

	// Default values
	s.apFast = 0.0
	s.apSlow = 0.0
	s.preT = -1000000000000000.0
	s.refractoryState = false
	s.refractoryCnt = 0
	s.efficacyTrace = 0.0
	s.psp = 0

	s.axon.Reset()
	s.dendrite.Reset()
}

// Integrate is the core the soma
func (s *Soma) Integrate(spanT, t int) (spike int) {
	dt := float64(t) - s.preT
	neuron := s.simJ.Neuron

	s.efficacyTrace = s.Efficacy(dt)

	// The dendrite will return a value that affects the soma.
	s.psp = s.dendrite.Integrate(spanT, t)

	// Default state
	s.axon.Reset()

	if s.refractoryState {
		// this algorithm should be the same as for the synapse or at least very
		// close.
		if s.refractoryCnt >= neuron.RefractoryPeriod {
			s.refractoryState = false
			s.refractoryCnt = 0
			// fmt.Printf("Refractory ended at (%d)\n", int(t))
		} else {
			s.refractoryCnt = s.refractoryCnt + 1
		}
	} else {
		if s.psp > s.threshold {
			// An action potential just occurred.
			// TODO Handle depolarization

			s.refractoryState = true

			// TODO
			// Generate a back propagating spike that fades spatial/temporally similar to CaDP model.
			// This spike affects forward in time.
			// The value is driven by the time delta of (preAPt - APt)
			s.axon.Set(1) // We set immediately because we are simulating a single neuron.
			// s.axon.Input(1)
			// fmt.Println(t)

			// Surge from action potential

			s.nFastSurge = s.APMax + s.apFast*neuron.FastSurge*math.Exp(-s.apFast/neuron.Tao)
			s.nSlowSurge = s.APMax + s.apSlow*neuron.SlowSurge*math.Exp(-s.apSlow/neuron.TaoS)

			// Reset time deltas
			s.preT = float64(t)
			dt = 0
		}
	}

	// Prior is for triplet
	s.apSlowPrior = s.apSlow

	// fmt.Printf("Soma:: %0.3f, %0.3f, psp:%0.3f\n", s.nFastSurge, s.nSlowSurge, s.psp)
	// println(soma.nFastSurge, ", ", soma.nSlowSurge, ", ", soma.ntao, ", ", soma.ntaoS)
	s.apFast = s.nFastSurge * math.Exp(-dt/neuron.Tao)
	s.apSlow = s.nSlowSurge * math.Exp(-dt/neuron.TaoS)

	// Collect this soma' values at this time step
	s.samples.CollectSoma(s, t)

	return s.axon.Output()
}

// Step after integration
func (s *Soma) Step() {
	s.axon.Step()
}

// ApSlowPrior ...
func (s *Soma) ApSlowPrior() float64 {
	return s.apSlowPrior
}

// EfficacyTrace soma's computed trace
func (s *Soma) EfficacyTrace() float64 {
	return s.efficacyTrace
}

// Efficacy based on TaoJ
func (s *Soma) Efficacy(dt float64) float64 {
	nMod := s.simJ.Neuron
	return 1.0 - math.Exp(-dt/nMod.TaoJ)
}

// =============================================================
// Sampling access
// =============================================================

// Output of soma's axon
func (s *Soma) Output() int {
	return s.axon.Output()
}

// APFast ...
func (s *Soma) APFast() float64 {
	return s.apFast
}

// APSlow ...
func (s *Soma) APSlow() float64 {
	return s.apSlow
}

// Psp ...
func (s *Soma) Psp() float64 {
	return s.psp
}
