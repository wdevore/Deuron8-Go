package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/datasamples"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

type environmentS struct {
	config          api.IModel
	sim             api.IModel
	synapsesPersist api.IModel

	samples api.ISamples
	// Synapses
	synapses []api.ISynapse

	stimulus          [][]int
	expandedStimulus  [][]int
	stimulusStreamCnt int

	relativePath string
	basePath     string

	simChan chan string
	cmd     string
	params  string

	simRunning bool

	// Runtime properties
	randomizerChoice             int // Global panel
	initialWeightValues          int // Global panel
	minRangeValue, maxRangeValue float64
	centerRangeValue             float64

	// Hard or Soft
	weightBounding   int
	softAcceleration float64
	softCurve        float64 // 1.0 = linear, 2.0 = parabola
}

// NewEnvironment ...
func NewEnvironment(relativePath, basePath string) api.IEnvironment {
	o := new(environmentS)
	o.relativePath = relativePath
	o.basePath = basePath

	o.minRangeValue = 5.0
	o.maxRangeValue = 10.0

	// load config.json
	o.loadProperties()

	config := o.config.Data().(*model.ConfigJSON)
	o.loadSynapses(config.SynapticPresets) // Ex: "synapses.json"

	o.samples = datasamples.NewSamples()

	o.loadStimulus()

	return o
}

func (e *environmentS) loadStimulus() {
	dataPath, err := filepath.Abs(e.relativePath)
	if err != nil {
		panic(err)
	}

	moData, _ := e.config.Data().(*model.ConfigJSON)

	eConfFile, err := os.Open(dataPath + e.basePath + moData.SourceStimulus + ".data")
	if err != nil {
		panic(err)
	}

	defer eConfFile.Close()

	e.expandedStimulus = [][]int{}

	scanner := bufio.NewScanner(eConfFile)
	for scanner.Scan() {
		pattern := scanner.Text()

		// Each line is the same length
		duration := len(pattern)

		// The array size is duration + (duration * StimExpander)
		// For example, if duration is 10 and stim_scaler is 3 then
		// size of stimulus is 10 + (10*3) = 40
		// StimExpander thus becomes an expanding factor. For every bit in
		// the pattern we append StimExpander 0s.
		if moData.StimulusScaler == 0 {
			// Special case of 0 then duration is unchanged (i.e. reflected)
			moData.StimulusScaler = 1 // Note: we don't call Changed() on purpose.
		} else {
			duration = (duration * moData.StimulusScaler)
		}

		expanded := make([]int, duration)
		stim := []int{}

		col := 0
		for _, c := range pattern {
			if c == '|' {
				expanded[col] = 1
				stim = append(stim, 1)
			} else {
				stim = append(stim, 0)
			}
			// Move col "past" the expanded positions.
			col += moData.StimulusScaler
		}

		e.stimulus = append(e.stimulus, stim)
		e.expandedStimulus = append(e.expandedStimulus, expanded)

		e.stimulusStreamCnt++
	}
}

func (e *environmentS) expandStimulus(scaler int) {
	// Reset expanded data
	e.expandedStimulus = [][]int{}

	// All channels are the same length, pick 0
	duration := (len(e.stimulus[0]) * scaler)

	// Iterate each channel and expand it.
	for _, stim := range e.stimulus {
		expanded := make([]int, duration)
		col := 0
		for _, spike := range stim {
			if spike == 1 {
				expanded[col] = 1
			}
			// Move col "past" the expanded positions.
			col += scaler
		}
		e.expandedStimulus = append(e.expandedStimulus, expanded)
	}
}

func (e *environmentS) loadProperties() {
	e.config = model.NewConfigModel(e.relativePath, e.basePath+"config.json")

	moData, ok := e.config.Data().(*model.ConfigJSON)

	if ok {
		fmt.Println("Config loaded.")
		fmt.Println("AutoSave: ", moData.AutoSave)
	} else {
		panic("Failed to load config.json")
	}

	// -------------------------------------------------------------
	// Load simulation property settings
	// -------------------------------------------------------------
	e.sim = model.NewSimModel(e.relativePath, e.basePath+"sim_model.json")

	simData, ok := e.sim.Data().(*model.SimJSON)

	if ok {
		fmt.Println("Default sim_model loaded.")
		fmt.Println("Synapses: ", simData.Synapses)
	} else {
		panic("Failed to load sim_model.json")
	}
}

func (e *environmentS) loadSynapses(file string) {
	e.synapsesPersist = model.NewSynapsePersist(e.relativePath, e.basePath+file)

	_, ok := e.synapsesPersist.Data().(*model.SynapsesJSON)

	if ok {
		fmt.Println("Synaptic data loaded.")
	} else {
		panic("Failed to load synapses.json")
	}
}

func (e *environmentS) Config() api.IModel {
	return e.config
}

func (e *environmentS) Sim() api.IModel {
	return e.sim
}

func (e *environmentS) SynapticModel() api.IModel {
	return e.synapsesPersist
}

func (e *environmentS) AddSynapse(synapse api.ISynapse) {
	e.synapses = append(e.synapses, synapse)
}

func (e *environmentS) Synapses() []api.ISynapse {
	return e.synapses
}

func (e *environmentS) Samples() api.ISamples {
	return e.samples
}

func (e *environmentS) StimulusCount() int {
	return e.stimulusStreamCnt
}

func (e *environmentS) Stimulus() [][]int {
	return e.expandedStimulus
}

func (e *environmentS) StimulusAt(idx int) []int {
	return e.expandedStimulus[idx]
}

func (e *environmentS) Run(state bool) {
	e.simRunning = state
}

func (e *environmentS) IsRunning() bool {
	return e.simRunning
}

// --------------------------------------------------------------
// Simple message command system
// --------------------------------------------------------------

// The issued "cmd" is recognized in Run.go's run() method
func (e *environmentS) IssueCmd(cmd string) {
	e.cmd = cmd

	switch cmd {
	case "propertyChange":
		switch e.params {
		case "StimulusScaler":
			fmt.Println("Scaler changed. Adjusting stimulus expansion.")
			moData, _ := e.config.Data().(*model.ConfigJSON)
			e.expandStimulus(moData.StimulusScaler)
		}
	}
}

func (e *environmentS) IsCmdIssued() bool {
	return e.cmd != ""
}

func (e *environmentS) CmdIssued() {
	e.cmd = ""
}

func (e *environmentS) Cmd() string {
	return e.cmd
}

func (e *environmentS) Parms() string {
	return e.params
}

func (e *environmentS) SetParms(parms string) {
	e.params = parms
}

// -------------------------------------------------------------
// Runtime properties
// -------------------------------------------------------------

func (e *environmentS) RandomizerField() int {
	return e.randomizerChoice
}

func (e *environmentS) SetRandomizerField(v int) {
	e.randomizerChoice = v
}

func (e *environmentS) InitialWeightValues() int {
	return e.initialWeightValues
}

func (e *environmentS) SetInitialWeightValues(v int) {
	e.initialWeightValues = v
}

func (e *environmentS) MinimumRangeValue() float64 {
	return e.minRangeValue
}

func (e *environmentS) SetMinimumRangeValue(v float64) {
	e.minRangeValue = v
}

func (e *environmentS) MaximumRangeValue() float64 {
	return e.maxRangeValue
}

func (e *environmentS) SetMaximumRangeValue(v float64) {
	e.maxRangeValue = v
}

func (e *environmentS) CenterRangeValue() float64 {
	return e.centerRangeValue
}

func (e *environmentS) SetCenterRangeValue(v float64) {
	e.centerRangeValue = v
}

func (e *environmentS) WeightBounding() int {
	return e.weightBounding
}

func (e *environmentS) SetWeightBounding(v int) {
	e.weightBounding = v
}
