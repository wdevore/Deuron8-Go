package datasamples

import (
	"fmt"
	"math"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

const (
	maxSynapses = 30
)

// Each synapse has one stream feeding into it.
// Some could be Noise while others are Stimulus

// SynapseSamples for graphs
type SynapseSamples struct {
	t      int
	id     int
	weight float64
	surge  float64
	psp    float64
	input  int // Stimulus
}

// T ...
func (y *SynapseSamples) T() int { return y.t }

// ID ...
func (y *SynapseSamples) ID() int { return y.id }

// Weight ...
func (y *SynapseSamples) Weight() float64 { return y.weight }

// Surge ...
func (y *SynapseSamples) Surge() float64 { return y.surge }

// Psp ...
func (y *SynapseSamples) Psp() float64 { return y.psp }

// Input ...
func (y *SynapseSamples) Input() int { return y.input }

// SomaSamples for graphs
type SomaSamples struct {
	t      int
	apFast float64
	apSlow float64
	output int
	psp    float64
}

// T ...
func (y *SomaSamples) T() int { return y.t }

// ApFast ...
func (y *SomaSamples) ApFast() float64 { return y.apFast }

// ApSlow ...
func (y *SomaSamples) ApSlow() float64 { return y.apSlow }

// Output ...
func (y *SomaSamples) Output() int { return y.output }

// Psp ...
func (y *SomaSamples) Psp() float64 { return y.psp }

type samples struct {
	// Synaptic data. There are N synapses and each is tracked
	// with their own collection.
	synData [][]api.ISynapseSample

	somaData []api.ISomaSample

	synapseSurgeMin, synapseSurgeMax   float64
	synapsePspMin, synapsePspMax       float64
	synapseWeightMin, synapseWeightMax float64

	somaPspMin, somaPspMax       float64
	somaAPFastMin, somaAPFastMax float64
	somaAPSlowMin, somaAPSlowMax float64
}

// NewSamples returns a samples collection
func NewSamples() api.ISamples {
	o := new(samples)

	o.Reset()

	return o
}

func (s *samples) Reset() {
	fmt.Println("Samples resetting.")
	s.synData = make([][]api.ISynapseSample, maxSynapses)
	s.somaData = []api.ISomaSample{}

	s.synapseSurgeMin = 1000000000000.0
	s.synapseSurgeMax = -1000000000000.0
	s.synapsePspMin = 1000000000000.0
	s.synapsePspMax = -1000000000000.0
	s.synapseWeightMin = 1000000000000.0
	s.synapseWeightMax = -1000000000000.0

	s.somaPspMin = 1000000000000.0
	s.somaPspMax = -1000000000000.0
	s.somaAPFastMin = 1000000000000.0
	s.somaAPFastMax = -1000000000000.0
	s.somaAPSlowMin = 1000000000000.0
	s.somaAPSlowMax = -1000000000000.0
}

func (s *samples) CollectSynapse(synapse api.ISynapse, id, t int) {
	// Check if a channel is already in play. Create a new channel if not.
	if s.synData[id] == nil {
		s.synData[id] = []api.ISynapseSample{}
	}

	s.synapseSurgeMin = math.Min(s.synapseSurgeMin, synapse.Surge())
	s.synapseSurgeMax = math.Max(s.synapseSurgeMax, synapse.Surge())
	s.synapsePspMin = math.Min(s.synapsePspMin, synapse.Psp())
	s.synapsePspMax = math.Max(s.synapsePspMax, synapse.Psp())
	s.synapseWeightMin = math.Min(s.synapseWeightMin, synapse.Weight())
	s.synapseWeightMax = math.Max(s.synapseWeightMax, synapse.Weight())
	// if s.synapseWeightMin == math.NaN() {
	// 	fmt.Println("oops")
	// }
	// if t > 140 && t < 160 {
	// 	fmt.Printf("(%d)%d:%0.3f:%0.3f ", id, t, s.synapseWeightMin, synapse.Weight())
	// }
	// fmt.Printf("(%03d:%03d) W:%0.3f, Sur:%0.3f, Psp:%0.3f, Min:%03.f, Max:%0.3f, I:%d,\n",
	// 	t, id, synapse.Weight(), synapse.Surge(), synapse.Psp(),
	// 	s.synapseWeightMin, s.synapseWeightMax,
	// 	synapse.Input(),
	// )

	s.synData[id] = append(s.synData[id],
		&SynapseSamples{
			t:      t,
			id:     synapse.ID(),
			weight: synapse.Weight(),
			surge:  synapse.Surge(),
			psp:    synapse.Psp(),
			// Input is either Noise or Stimulus
			input: synapse.Input(),
		},
	)
}

func (s *samples) CollectSoma(soma api.ISoma, t int) {
	s.somaPspMin = math.Min(s.somaPspMin, soma.Psp())
	s.somaPspMax = math.Max(s.somaPspMax, soma.Psp())
	s.somaAPFastMin = math.Min(s.somaAPFastMin, soma.APFast())
	s.somaAPFastMax = math.Max(s.somaAPFastMax, soma.APFast())
	s.somaAPSlowMin = math.Min(s.somaAPSlowMin, soma.APSlow())
	s.somaAPSlowMax = math.Max(s.somaAPSlowMax, soma.APSlow())

	s.somaData = append(s.somaData,
		&SomaSamples{
			t:      t,
			apFast: soma.APFast(),
			apSlow: soma.APSlow(),
			psp:    soma.Psp(),
			output: soma.Output(), // Spikes
		},
	)
}

// func (s *samples) SynapticData() map[int][]api.ISynapseSample {
// 	return nil //s.synData
// }

func (s *samples) SynapticData() [][]api.ISynapseSample {
	return s.synData
}

func (s *samples) SynapseData(id int) []api.ISynapseSample {
	return s.synData[id]
}

func (s *samples) SomaData() []api.ISomaSample {
	return s.somaData
}

func (s *samples) CollectDendrite(dendrite api.IDendrite, t int) {

}

func (s *samples) SynapseSurgeMin() float64 {
	return s.synapseSurgeMin
}

func (s *samples) SynapseSurgeMax() float64 {
	return s.synapseSurgeMax
}

func (s *samples) SynapsePspMin() float64 {
	return s.synapseSurgeMin
}

func (s *samples) SynapsePspMax() float64 {
	return s.synapseSurgeMax
}

func (s *samples) SynapseWeightMin() float64 {
	return s.synapseWeightMin
}

func (s *samples) SynapseWeightMax() float64 {
	return s.synapseWeightMax
}

func (s *samples) SomaPspMin() float64 {
	return s.somaPspMin
}

func (s *samples) SomaPspMax() float64 {
	return s.somaPspMax
}

func (s *samples) SomaAPFastMin() float64 {
	return s.somaAPFastMin
}

func (s *samples) SomaAPFastMax() float64 {
	return s.somaAPFastMax
}

func (s *samples) SomaAPSlowMin() float64 {
	return s.somaAPSlowMin
}

func (s *samples) SomaAPSlowMax() float64 {
	return s.somaAPSlowMax
}
