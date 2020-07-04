package api

// ISoma is main body of the cell as an abstraction.
type ISoma interface {
	Initialize()
	Reset()
	Integrate(spanT, t int) (psp float64)

	APFast() float64
	ApSlowPrior() float64
	EfficacyTrace() float64

	Output() int
}
