package model

// SynapseDefaultsJSON is default values for a synapses
type SynapseDefaultsJSON struct {
	TaoP             float64
	TaoN             float64
	Mu               float64
	Distance         float64
	Lambda           float64
	Amb              float64
	W                float64
	Alpha            float64
	LearningRateSlow float64
	LearningRateFast float64
	TaoI             float64
	Ama              float64
}

// CompartmentJSON is compartment persisted data
type CompartmentJSON struct {
	WeightMax       float64
	ID              int
	WeightDivisor   float64
	SynapseDefaults *SynapseDefaultsJSON
}

// DendriteJSON is dendrite persisted data
type DendriteJSON struct {
	Length       float64
	TaoEff       float64
	MinPSPValue  float64
	ID           int
	Compartments []*CompartmentJSON
}

// NeuronJSON is neuron persisted data
type NeuronJSON struct {
	Tao              float64
	FastSurge        float64
	WMax             float64
	TaoJ             float64
	WMin             float64
	ID               int
	TaoS             float64
	APMax            float64
	Threshold        float64
	RefractoryPeriod float64
	SlowSurge        float64
	Dendrites        *DendriteJSON
}

// SimJSON is simulation persisted data
type SimJSON struct {
	Synapses                    int
	ActiveSynapse               int
	PoissonPatternSpread        int
	PercentOfExcititorySynapses float64
	// If Hertz = 0 then stimulus is distributed as poisson.
	// Hertz is = cycles per second (or 1000ms per second)
	// 10Hz = 10 applied in 1000ms or every 100ms = 1000/10Hz
	// This means a stimulus is generated every 100ms which also means the
	// Inter-spike-interval (ISI) is fixed at 100ms
	Hertz             int
	PoissonPatternMax float64
	PoissonPatternMin float64
	StimulusScaler    float64

	// Poisson stream Lambda
	// Firing rate = spikes over an interval of time or
	// Poisson events per interval of time.
	// For example, spikes in a 1 sec span.
	NoiseLambda float64
	NoiseCount  int

	Neuron *NeuronJSON
}
