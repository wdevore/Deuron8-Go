package datasamples

import "github.com/wdevore/Deuron8-Go/neuron_simulation/api"

// SynapseSamples for graphs
type SynapseSamples struct {
	t      int
	id     int
	weight float64
	surge  float64
	psp    float64
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
	// Synapse data
	synData  []api.ISynapseSample
	somaData []api.ISomaSample
}

// NewSamples returns a samples collection
func NewSamples() api.ISamples {
	o := new(samples)

	o.Reset()

	return o
}

func (s *samples) Reset() {
	s.synData = []api.ISynapseSample{}
	s.somaData = []api.ISomaSample{}
}

func (s *samples) CollectSynapse(synapse api.ISynapse, t int) {
	s.synData = append(s.synData,
		&SynapseSamples{
			t:      t,
			id:     synapse.ID(),
			weight: synapse.Weight(),
			surge:  synapse.Surge(),
			psp:    synapse.Psp(),
		},
	)
}

func (s *samples) CollectDendrite(dendrite api.IDendrite, t int) {

}

func (s *samples) CollectSoma(soma api.ISoma, t int) {
	s.somaData = append(s.somaData,
		&SomaSamples{
			t:      t,
			apFast: soma.APFast(),
			apSlow: soma.APSlow(),
			output: soma.Output(), // Spikes
			psp:    soma.Psp(),
		},
	)
}

func (s *samples) SynapseData() []api.ISynapseSample {
	return s.synData
}

func (s *samples) SomaData() []api.ISomaSample {
	return s.somaData
}
