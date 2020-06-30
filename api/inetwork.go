package api

// INetwork represents the network of neurons under simulation.
type INetwork interface {

	// Step processes the current time tick.
	// The output is transfered to the input of a synapse for the
	// next time step.

	// Next moves to the next state

	// Load from either a previous simulation or start new.
	// Use IConfig to determine if we need to load a previous
	// construction or start new
	Load(IConfig)

	// Save the network's construction and configuration.
	Save()
}
