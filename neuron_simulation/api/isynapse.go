package api

// ISynapse is part of a compartment and dendrite
type ISynapse interface {
	Initialize()
	Reset()

	SetStream(IBitStream)
	Input() int   // Output from stream "into" this synapse
	SetType(bool) // Inhibit=false, excititory=true

	PreIntegrate()
	Integrate(spanT, t int) (value, w float64)
	PostIntegrate()

	// Data sample fields
	ID() int
	SetID(int)

	InitialWeight() float64
	SetInitialWeight(float64)
	Weight() float64
	SetWeight(float64)

	Surge() float64
	Psp() float64

	SetToDefaults()
}
