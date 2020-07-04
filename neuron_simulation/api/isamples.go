package api

// ISamples represents samples taken at each time step
type ISamples interface {
	CollectSynapse(synapse ISynapse, t int)
	CollectDendrite(dendrite IDendrite, t int)
	CollectSoma(soma ISoma, t int)
}
