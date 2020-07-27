# Neuron simulation
The code in this folder is a back port of the [Julia single neuron](https://github.com/wdevore/Deuron7-Julia) simulation to GoLang.

# Simulation
The simulation tests the functionality of a single neuron and what range of parameters values the neuron is functional within.

The parameters are tuned against a given number of synapses. As the synapse count changes the neuron will need re-tuning.

# Folders
* neuron_simulation:
    * api: Interfaces
    * app: The single neuron simulation driven and view with a GUI
    * cell: The support objects making up the simulation.
    * cellsimulation: the go coroutine that exercises the objects in "cell" folder
    * data: the data that configures and sources the simulation
    * datasamples: the sampling of the sim output
    * graphs: the IMGUI graphs for the data
    * misc: guess what?
    * model: the JSON code for handling the .json files
    * streams: the synaptic input streams
    * tests: guess what?