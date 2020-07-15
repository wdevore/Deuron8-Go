package model

// SynapseJSON is synapse persisted data
type SynapseJSON struct {
	TaoP             float64
	TaoN             float64
	Mu               float64
	Distance         float64
	ID               int
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
	WeightMax     float64
	ID            int
	WeightDivisor float64
	Synapses      []*SynapseJSON
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
	Hertz                       float64
	FiringRate                  float64
	PoissonPatternMax           float64
	PoissonPatternMin           float64
	StimulusScaler              float64
	Neuron                      *NeuronJSON
}
