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
	config   api.IModel
	sim      api.IModel
	samples  api.ISamples
	stimulus [][]int

	relativePath string
	basePath     string

	autostop bool

	simChan chan string
	cmd     string
}

// NewEnvironment ...
func NewEnvironment(relativePath, basePath string) api.IEnvironment {
	o := new(environmentS)
	o.relativePath = relativePath
	o.basePath = basePath

	o.loadProperties()
	o.autostop = false
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

	e.stimulus = [][]int{}

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
		if moData.StimExpander == 0 {
			// Special case of 0 then duration is unchanged (i.e. reflected)
			moData.StimExpander = 1 // Note: we don't call Changed() on purpose.
		} else {
			duration = (duration * moData.StimExpander)
		}

		expanded := make([]int, duration)

		col := 0
		for _, c := range pattern {
			if c == '|' {
				expanded[col] = 1
			}
			// Move col "past" the expanded positions.
			col += moData.StimExpander
		}

		e.stimulus = append(e.stimulus, expanded)
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

func (e *environmentS) Config() api.IModel {
	return e.config
}

func (e *environmentS) Sim() api.IModel {
	return e.sim
}

func (e *environmentS) Samples() api.ISamples {
	return e.samples
}

func (e *environmentS) Stimulus() [][]int {
	return e.stimulus
}

func (e *environmentS) StimulusAt(idx int) []int {
	return e.stimulus[idx]
}

func (e *environmentS) AutoStop(auto bool) {
	e.autostop = auto
}

func (e *environmentS) IsAutoStop() bool {
	return e.autostop
}

func (e *environmentS) IssueCmd(cmd string) {
	e.cmd = cmd
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
