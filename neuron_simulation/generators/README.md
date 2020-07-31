# Generators
The application in this folder generates synaptic presets for the simulator.

# Input
A json file controls the generator.

Each property of the synapse can be *generated* via a set of controls:

* "Enabled": true/false  // If disabled then the center value is used other wise the Min/Max values combined with Random Lerp is used.
* "Min": float64 // The minimum random value that be generated
* "Max": float64 // The maximum value
* "Center": float64 // The preset value or **center** around which values are created.

The *Center* value's position relative to Min/Max determines how likely a value is generated above or below

```
   Min                      Center                   Max
    |-------------------------|-----------------------|

or

   Min     Center                                    Max
    |--------|----------------------------------------|

```

In the second range more Max values are generated than Min values.