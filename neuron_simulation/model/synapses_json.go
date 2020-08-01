package model

// SynapseJSON is synapse persisted data
type SynapseJSON struct {
	Excititory       bool
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

// SynapsesJSON is synapse state persistance
type SynapsesJSON struct {
	Synapses []*SynapseJSON
}
