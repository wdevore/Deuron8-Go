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
* Add initial weight line to weight graph

When the GUI is launched the simulator is built and initialized with Presets.

At the start of a simulation the model is configured accorded to the initializer choice. The model is then copied to the runtime and the simulation is started. Once the simulation is complete the runtime is copied back to the model. The model can then be saved to **synapses.json**

Types of initializations possible:
* *Preset*: the synaptic properties are reset back to a set of Presets values from a fixed **synapsed_presets_*N*.json**.
* *Current*: the synaptic properties are not changed but continue forward. Any changes made to the model are used in the next simulation.
* *Random*: the synaptic properties are randomly changed prior to simulation. The only property currently supported is Weight.

### Generators
**Done** -- A generator program will create json files for synapses. It reads a configuration json file to drive the generator. The generated file can be used by the simulator as a preset.

## Notes:
The weights are typically the short term memory until learning **Effort** is increased.
The other parameters, for example Tao, are controlled by a Meta system. They are the slower changing parms that represent the charateristics of the neuron but also represent information at a meta level.

We have two json files: A Preset that can't be saved over and a version that is overlayed at the end of a simulation.

add random property generators, for example, randomly generate weights centered around a fixed value with a mean variance. Zero is a typical center value.
maybe have a listbox that you choose a property then select a max variance and center value. There could be presets.

Each time the sim starts we check if we use the fixed defaults, the current
weights or random values.