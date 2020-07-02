package api

// ISynapse is part of a compartment and dendrite
type ISynapse interface {
	Initialize()
	Reset()
	SetStream(IBitStream)
	SetType(bool) // Inhibit=false, excititory=true

	PreIntegrate()
	Integrate()
	PostIntegrate()

	TripleIntegration(spanT, t int64)
	Efficacy(dt float64)
	WeightFactor(potentiation bool)
}
