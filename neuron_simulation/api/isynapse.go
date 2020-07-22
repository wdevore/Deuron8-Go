package api

// ISynapse is part of a compartment and dendrite
type ISynapse interface {
	Initialize(bool)
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

	Weight() float64
	Surge() float64
	Psp() float64
}
