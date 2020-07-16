package api

// ISamples represents samples taken at each time step
type ISamples interface {
	CollectSynapse(synapse ISynapse, t int)
	CollectDendrite(dendrite IDendrite, t int)
	CollectSoma(soma ISoma, t int)

	Reset()

	SynapseData() []ISynapseSample
	SomaData() []ISomaSample
}

// ISynapseSample ...
type ISynapseSample interface {
	T() int
	ID() int
	Weight() float64
	Surge() float64
	Psp() float64
}

// ISomaSample ...
type ISomaSample interface {
	T() int
	ApFast() float64
	ApSlow() float64
	Output() int
	Psp() float64
}
