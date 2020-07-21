package cellsimulation

import (
	"fmt"
	"time"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/cell"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/streams"
)

// Simulator ...
type Simulator struct {
	environment api.IEnvironment

	// Noise streams
	noises []api.IBitStream

	// Stimulus
	stimuli []api.IBitStream

	neuron api.ICell
	t      int

	running bool
	step    bool
}

// NewSimulator create simulation object
func NewSimulator(environment api.IEnvironment) *Simulator {
	o := new(Simulator)
	o.environment = environment
	return o
}

// Build cell system
func (si *Simulator) Build() {
	simMod := si.environment.Sim()
	siData, _ := simMod.Data().(*model.SimJSON)
	moData, _ := si.environment.Config().Data().(*model.ConfigJSON)

	// First we create the Noise (Poisson) streams. Each stream will
	// be routed to a unique synapse. We need a collection of them so
	// we can exercise them on each simulation step.
	seed := uint64(5000)
	for i := 0; i < siData.NoiseCount; i++ {
		noise := streams.NewPoissonStream(seed, 4.0)
		si.noises = append(si.noises, noise)
		seed += 5000
	}
	fmt.Println("Poisson Noise streams created")

	// Now create the stimulus streams
	for i := 0; i < si.environment.StimulusCount(); i++ {
		stimAry := si.environment.StimulusAt(i)
		stim := streams.NewStimulusStream(stimAry, siData.Hertz)
		si.stimuli = append(si.stimuli, stim)
	}

	fmt.Println("Stimulus streams created")

	samples := si.environment.Samples()

	// Create cell dependencies starting with soma first.
	soma := cell.NewSoma(simMod, samples)

	dendrite := cell.NewDendrite(simMod, soma)
	soma.SetDendrite(dendrite)
	compartment := cell.NewCompartment(simMod, dendrite, soma)
	dendrite.AddCompartment(compartment)

	axon := cell.NewAxonZeroDelay()
	soma.SetAxon(axon)

	// We need a synapse for each stream, both Noise and Stimulus
	// Noise first:
	for _, noise := range si.noises {
		synapse := cell.NewSynapse(simMod, samples, soma, dendrite, compartment)
		// route noise to synapse
		synapse.SetStream(noise)
		compartment.AddSynapse(synapse)
	}

	// Now stimulus:
	for _, stimulus := range si.stimuli {
		synapse := cell.NewSynapse(simMod, samples, soma, dendrite, compartment)
		// route noise to synapse
		synapse.SetStream(stimulus)
		compartment.AddSynapse(synapse)
	}

	// Now create the single Neuron that this simulation execises
	si.neuron = cell.NewCell(simMod, soma)
	si.neuron.Initialize()

	fmt.Println("Simulation duration: ", moData.Duration)
	fmt.Println("Simulation built")
}

// Run simulation
func (si *Simulator) Run(ch chan string) {
	loop := true
	started := false

	moData, _ := si.environment.Config().Data().(*model.ConfigJSON)
	duration := moData.Duration

	for loop {
		select {
		case cmd := <-ch:
			switch cmd {
			case "stop":
				fmt.Println("Stopping simulation...")
				si.running = false
			case "run":
				fmt.Println("Running simulation...")
				si.reset()
				si.running = true
				si.step = false
			case "step":
				fmt.Println("Stepping simulation at t=", si.t)
				si.step = true
				si.running = false
			case "reset":
				fmt.Println("Simulation has been reset")
				si.reset()
				si.running = false
			case "killSim":
				loop = false
			}
		default:
			if si.running {
				// Running means running completely through a simulation
				// for the specified duration (in milliseconds).
				complete := si.simulate(si.t, duration)
				if complete {
					si.reset()
					si.running = false
					fmt.Println("Simulation stopped")
				}
				si.t++
			} else if si.step {
				complete := si.simulate(si.t, duration)
				if complete {
					si.reset()
				}
				si.t++
				si.step = false
			}
		}

		if !started {
			started = true
			fmt.Println("Simulator is ready.")
		}

		<-time.After(time.Millisecond * 33)
	}

	fmt.Println("Simulation coroutine exited")
}

// simulate takes a single simulation step.
func (si *Simulator) simulate(t int, duration int) bool {
	si.neuron.Integrate(0, t)

	// Step all streams. This causes each stream to update and move
	// its internal value to its output for the next integration.
	for _, noise := range si.noises {
		noise.Step()
	}

	for _, stimulus := range si.stimuli {
		stimulus.Step()
	}

	si.neuron.Step()

	return t > duration
}

func (si *Simulator) reset() {
	si.t = 0

	si.neuron.Reset()

	for _, noise := range si.noises {
		noise.Reset()
	}

	for _, stimulus := range si.stimuli {
		stimulus.Reset()
	}
}
