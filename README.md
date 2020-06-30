# Deuron 8 written in Go
The simulation for a single neuron is still written in *Julia*.

# Simulation
The simulation makes two "passes" through the network for each
time step. Each step equals 10? microseconds.

Each simulation step works on the previous input. Each component
maintains an internal *current-state* and *next-state*.

The first pass exercises the inputs on each neuron.

The second pass moves the results of the first pass to the neuron's output. Note: the output of one neuron is the input of another, except for stimulus.

