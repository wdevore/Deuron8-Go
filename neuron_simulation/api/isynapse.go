package api

// ISynapse is part of a compartment and dendrite
type ISynapse interface {
	Initialize()
	Reset()
	SetStream(IBitStream)
	SetType(bool) // Inhibit=false, excititory=true

	PreIntegrate()
	Integrate(spanT, t int) (value, w float64)
	PostIntegrate()
}
