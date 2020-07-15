package api

// ICell is an ecapsulation
type ICell interface {
	Reset()
	Integrate(spanT, t int) (spike int) // Feeds into an Axon
	Output() int
}
