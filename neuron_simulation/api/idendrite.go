package api

type IDendrite interface {
	Initialize()
	APEfficacy(distance float64) float64
	AddCompartment(ICompartment)
}
