package stimulus

// Scanline represents a connection between pixels
// and neurons. Where neurons are mostly Sensory bi-polar types.
//
// Each pixel is expanded into binary bit slices.
// x, y are 10bits (when dealing with 1024x1024 image),
// and color components and Intensity are 8bits.
//
// As the sim runs the stimulus component scans the image line
// by line. This "line" feeds into and impacts a layer of neurons
// called the Sensory neurons.
//
// During configuration the scanline buffer's bit positions are
// either randomly connected to various synapses or
// reconstructed from a previous configuration.
// The connections are synapses within the Edge layer.
type Scanline struct {
	// Neuron Ids. Each Id represents a connection to a synapse in the
	// Edge input layer.
}

// Step process the current time tick.
// The output is transfered to the input of a synapse for the
// next time step.

// Next moves to the next scanline in image

// Configure from either a previous simulation or start new.

// Persist the Scanline's connection neuron Ids.
