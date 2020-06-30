# Network
The network is constructed using a radius/radial/density random approach where each neuron makes between N and M connections with its neighbors. This is a sparely connected network.

The network is two 2D grids. The neuron (Principle Cells) grid 1000x1000 in size. The IN grid is ~30% smaller than the neuron grid. An IN's axon will typically connect between 1 and 5 local PCs.

Over time, if the EAs can successfully do their job, Attractor points (aka Knowledge Points **KP**s) form. The EA can adjust both weights and connections in order to form KP fields.

Networks are persisted to binary files.
Each neuron is stored along with its properties.
During storage the neuron's x,y position is stored. Which is used to load the neuron.

# Neuron
The Neuron is the core of the simulation, and is based on Deuron7.

## Neuron Properties
* Grid position
* Sim properties, example Tau
* 

## Sampling
Sampling is the data that is collected during a Step

---

# Visuals
Several visual animations are created:

* Spike graphs. A plain graph that shows when spikes occur overtime. A rate graph showing colors based on firing rates, white->red.
* Connectivity graphs