package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

// SynapsesPersist is synapse presistance data
type SynapsesPersist struct {
	Synapses SynapsesJSON

	relativePath string
	file         string

	changed bool
}

// NewSynapsePersist creates and load data
func NewSynapsePersist(relativePath, file string) api.IModel {
	o := new(SynapsesPersist)

	o.Synapses = SynapsesJSON{}

	o.relativePath = relativePath
	o.file = file

	o.Load()

	return o
}

// Data returns the json loaded app data
func (m *SynapsesPersist) Data() interface{} {
	return &m.Synapses
}

// SetActiveSynapse not used
func (m *SynapsesPersist) SetActiveSynapse(id int) {
}

// Changed marks model dirty
func (m *SynapsesPersist) Changed() {
	m.changed = true
}

// Clean marks model NOT-dirty
func (m *SynapsesPersist) Clean() {
	m.changed = false
}

// Load model from disk
func (m *SynapsesPersist) Load() {
	dataPath, err := filepath.Abs(m.relativePath)
	if err != nil {
		panic(err)
	}

	eFile, err := os.Open(dataPath + m.file)
	if err != nil {
		panic(err)
	}

	defer eFile.Close()

	bytes, err := ioutil.ReadAll(eFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &m.Synapses)
	if err != nil {
		panic(err)
	}
}

// Save model to disk
func (m *SynapsesPersist) Save() {
	if m.changed {
		fmt.Println("Saving synaptic data...")
		indentedJSON, _ := json.MarshalIndent(m.Synapses, "", "  ")

		dataPath, err := filepath.Abs(m.relativePath)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(dataPath+m.file, indentedJSON, 0644)
		if err != nil {
			log.Fatalln("ERROR:", err)
		}

		m.Clean()
		fmt.Println("Synaptic data saved")
	}
}
