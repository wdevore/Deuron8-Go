# Simulation
The simulation is a Step consisting of two passes.

The simulation generates several different outputs which are mostly images sequences. The images are then combined into a lossless movie format of type *H.264* via Blender or FFMpeg.

## First pass
This pass operates on the current (internal state and input values).

## Second pass
This pass moves the internal state to each neuron's output
which makes it available for the **First** pass.

# Samples
The goal of sampling is to collect data in order to form animations.

There several animations that are created:
* Spike events: This animation shows neurons spiking over time
* PSP decay: This animation shows colors of a neuron's PSP value over time.

We also need to show the flow of "info spikes". How do we do this?
We could draw lines from Pre to Post. I think this means we need to keep previous state as well. This would allow us to track progression from previous to current states.

If a neuron spikes as a direct reaction to the input spikes present at the time of spiking then this means we need to track where the input spike came from. This would allow us to draw lines from pre to post and should give a form of flow.


# Viewers
There is currently only one type of interactive viewer in design and that is a viewer that allows moving around the image to view connectivity.

The other type of view is simply a movie viewer.