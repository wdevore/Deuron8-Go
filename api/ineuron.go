package api

// INeuron represents a single neuron.
type INeuron interface {

	// Integrate processes the current time tick.
	// This is considered the First pass.
	Integrate() float64

	// Next transfers the internal state to the output.
	// This is considered the Second pass.
	Next()

	// Load from either a previous simulation or start new.
	// Use IConfig to determine if we need to load a previous
	// construction or start new
	Load(IConfig)

	// Save the construction and configuration.
	Save()
}
