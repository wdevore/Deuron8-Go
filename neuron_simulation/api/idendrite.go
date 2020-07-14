package api

type IDendrite interface {
	Initialize()
	Reset()
	Integrate(spanT, t int) (psp float64)
	APEfficacy(distance float64) float64
	AddCompartment(ICompartment)
}
