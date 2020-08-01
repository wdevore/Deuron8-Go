package cellsimulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/cell"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/misc"
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

	// -----------------------------------------------------------------
	// First we create the Noise (Poisson) streams. Each stream will
	// be routed to a unique synapse. We need a collection of them so
	// we can exercise them on each simulation step.
	seed := uint64(5000)
	for i := 0; i < siData.NoiseCount; i++ {
		noise := streams.NewPoissonStream(seed, siData.NoiseLambda)
		si.noises = append(si.noises, noise)
		seed += 5000
	}
	fmt.Println("Poisson Noise streams created")

	// -----------------------------------------------------------------
	// Now create the stimulus streams
	for i := 0; i < si.environment.StimulusCount(); i++ {
		stimAry := si.environment.StimulusAt(i)
		stim := streams.NewStimulusStream(stimAry, siData.Hertz)
		si.stimuli = append(si.stimuli, stim)
	}

	fmt.Println("Stimulus streams created")

	samples := si.environment.Samples()

	// -----------------------------------------------------------------
	// Create cell dependencies starting with soma first.
	soma := cell.NewSoma(simMod, samples)

	dendrite := cell.NewDendrite(simMod, soma)
	soma.SetDendrite(dendrite)
	compartment := cell.NewCompartment(simMod, dendrite, soma)

	axon := cell.NewAxonZeroDelay()
	soma.SetAxon(axon)

	genSynID := 0

	// -----------------------------------------------------------------
	// We need a synapse for each stream, both Noise and Stimulus
	for _, stimulus := range si.stimuli {
		synapse := cell.NewSynapse(si.environment,
			soma, dendrite, compartment, genSynID)
		// route noise to synapse
		synapse.SetStream(stimulus)
		synapse.Initialize()
		si.environment.AddSynapse(synapse)
		genSynID++
	}

	// -----------------------------------------------------------------
	for _, noise := range si.noises {
		synapse := cell.NewSynapse(si.environment,
			soma, dendrite, compartment, genSynID)
		// route noise to synapse
		synapse.SetStream(noise)
		synapse.Initialize()
		si.environment.AddSynapse(synapse)
		genSynID++
	}

	// -----------------------------------------------------------------
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
				si.environment.Run(si.running)
				si.step = false
			case "step":
				fmt.Println("Stepping simulation at t=", si.t)
				si.step = true
				si.running = false
			case "reset":
				fmt.Println("Simulation has been reset")
				si.reset()
				si.running = false
			case "propertyChange":
				si.propertyChange(si.environment)
			case "killSim":
				loop = false
			}
		default:
			duration := moData.Duration
			if si.running {
				// Running means running completely through a simulation
				// for the specified duration (in milliseconds).
				complete := si.simulate(si.t, duration)
				if complete {
					si.running = false
					si.environment.Run(si.running)
					fmt.Println("Simulation finished for duration of: ", duration)
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

		if !si.running {
			<-time.After(time.Millisecond * 33)
		}
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

	si.environment.Samples().Reset()

	// Set initial values for each synapse: Presets, Current or Random?
	switch si.environment.InitialWeightValues() {
	case 0: // Presets
		si.synapsePresets()
	case 2: // Random
		si.synapseWeightRandomizer()
	}

	if si.environment.InitialWeightValues() != 1 {
		// Now transfer model to synapses
		for _, syn := range si.environment.Synapses() {
			syn.Initialize()
		}
	}

}

func (si *Simulator) synapsePresets() {
	si.environment.SynapticModel().Load()

	synsJ, _ := si.environment.SynapticModel().Data().(*model.SynapsesJSON)

	simMod, _ := si.environment.Sim().Data().(*model.SimJSON)
	defs := simMod.Neuron.Dendrites.Compartments[0].SynapseDefaults

	// Modify the model to reflect source
	for _, syn := range synsJ.Synapses {
		syn.TaoP = defs.TaoP
		syn.TaoN = defs.TaoN
		syn.Mu = defs.Mu
		syn.Distance = defs.Distance
		syn.Lambda = defs.Lambda
		syn.Amb = defs.Amb
		syn.W = defs.W
		syn.Alpha = defs.Alpha
		syn.LearningRateSlow = defs.LearningRateSlow
		syn.LearningRateFast = defs.LearningRateFast
		syn.TaoI = defs.TaoI
		syn.Ama = defs.Ama
	}

}

func (si *Simulator) synapseWeightRandomizer() {
	// parms := strings.Split(si.environment.Parms(), ",")

	// switch parms[0] {
	// case "Weight":
	synapses, _ := si.environment.SynapticModel().Data().(*model.SynapsesJSON)

	// Use the Min/Max values to bound the Lerp
	min := si.environment.MinimumRangeValue()
	max := si.environment.MaximumRangeValue()
	center := si.environment.CenterRangeValue()

	for _, syn := range synapses.Synapses {
		l := misc.Linear(min, max, center)
		r := rand.Float64()

		if r > l {
			// Center -> Max wins
			syn.W = misc.Lerp(center, max, rand.Float64())
			continue
		}

		// Min -> Center wins
		syn.W = misc.Lerp(min, center, rand.Float64())
	}
	// }

}

func (si *Simulator) propertyChange(environment api.IEnvironment) {
	simMod := environment.Sim()

	switch environment.Parms() {
	case "PoissonLambda":
		// Update all Noise streams
		for _, noise := range si.noises {
			noise.Update(simMod)
		}
	case "StimulusScaler":
		// The Environment has already scaled the stimulus
		// prior to the propertyChange event here.
		siData, _ := simMod.Data().(*model.SimJSON)
		si.stimuli = nil

		// We need to reassign the same stream to the same synapse.
		for i, bitstream := range environment.Stimulus() {
			// Rebind all patterns to each stimulus stream
			synapse := si.environment.Synapses()[i]
			stim := streams.NewStimulusStream(bitstream, siData.Hertz)
			synapse.SetStream(stim)
			si.stimuli = append(si.stimuli, stim)
		}
	case "Duration":
		// Currently not used
	}
}
