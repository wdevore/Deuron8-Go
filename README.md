# Deuron 8 written in Go
The simulation for a single neuron is still written in *Julia*.

# Simulation
The simulation makes two "passes" through the network for each
time step. Each step equals 10? microseconds.

Each simulation step works on the previous input. Each component
maintains an internal *current-state* and *next-state*.

The first pass exercises the inputs on each neuron.

The second pass moves the results of the first pass to the neuron's output. Note: the output of one neuron is the input of another, except for stimulus.

# Folders
* Deuron8-Go
    * api: the main api for the network approach
    * config:
    * examples:
    * log:
    * neuron_simulation: see folder [readme.md](neuron_simulation/readme.md)
    * simulation:
    * tests:
    * tools:

# Tasks
## Synaptic properties
* Add save/load for synapse properties
* Randomize specific properties, for example, weights
* Choose to use defaults or random

We need to read each runtime synapse and add/update json property
The save synapse properties

add random property generators, for example, randomly generate weights centered around a fixed value with a mean variance. Zero is a typical center value.
maybe have a listbox that you choose a property then select a max variance and center value. There could be presets.

-- need bool "Initial weights" verse randomally generated. what sources the initial weight values used at the start of a simulation? config indicates it.

Each time the sim starts we check if we use the fixed defaults, the current
weights or random values.