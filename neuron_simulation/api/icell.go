package api

// ICell is an ecapsulation
type ICell interface {
	Initialize()
	Reset()
	Integrate(spanT, t int) (spike int) // Feeds into an Axon
	Step()
	Output() int
}
