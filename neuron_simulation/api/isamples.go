package api

// ISamples represents samples taken at each time step
type ISamples interface {
	CollectSynapse(synapse ISynapse, id, t int)
	CollectDendrite(dendrite IDendrite, t int)
	CollectSoma(soma ISoma, t int)

	Reset()

	// Synaptic data
	SynapticData() [][]ISynapseSample
	SynapseData(int) []ISynapseSample

	SynapseSurgeMin() float64
	SynapseSurgeMax() float64

	SynapsePspMin() float64
	SynapsePspMax() float64

	SynapseWeightMin() float64
	SynapseWeightMax() float64

	// Soma data
	SomaData() []ISomaSample
	SomaPspMin() float64
	SomaPspMax() float64
	SomaAPFastMin() float64
	SomaAPFastMax() float64
	SomaAPSlowMin() float64
	SomaAPSlowMax() float64
}

// ISynapseSample one for each synapse
type ISynapseSample interface {
	T() int
	ID() int
	Weight() float64
	Surge() float64
	Psp() float64
	Input() int
}

// ISomaSample for a single soma
type ISomaSample interface {
	T() int
	ApFast() float64
	ApSlow() float64
	Output() int
	Psp() float64
}
