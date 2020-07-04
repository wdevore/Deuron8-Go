package api

// ICell is an ecapsulation
type ICell interface {
	Reset()
	Integrate(spanT, t int) (psp float64) // Feeds into an Axon
	Output() int
}
