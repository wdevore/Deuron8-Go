package api

// ICompartment is a part of a cell
type ICompartment interface {
	Initialize()
	Reset()
	AddSynapse(ISynapse)
	Integrate(spanT, t int) (psp, totalWeight float64)
}
