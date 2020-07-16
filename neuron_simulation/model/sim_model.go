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

// SimModel is the simulation presistance data
type SimModel struct {
	Sim SimJSON

	relativePath string
	file         string

	changed bool
}

// NewSimModel creates and loads sim data
func NewSimModel(relativePath, file string) api.IModel {
	o := new(SimModel)

	o.Sim = SimJSON{}

	o.relativePath = relativePath
	o.file = file

	o.Load()

	return o
}

// Data returns the json loaded app data
func (m *SimModel) Data() interface{} {
	return &m.Sim
}

// SetActiveSynapse ...
func (m *SimModel) SetActiveSynapse(id int) {
	m.Sim.ActiveSynapse = id
	m.changed = true
}

// Changed marks model dirty
func (m *SimModel) Changed() {
	m.changed = true
}

// Clean marks model NOT-dirty
func (m *SimModel) Clean() {
	m.changed = false
}

// Load model from disk
func (m *SimModel) Load() {
	dataPath, err := filepath.Abs(m.relativePath)
	if err != nil {
		panic(err)
	}

	eConfFile, err := os.Open(dataPath + m.file)
	if err != nil {
		panic(err)
	}

	defer eConfFile.Close()

	bytes, err := ioutil.ReadAll(eConfFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &m.Sim)
	if err != nil {
		panic(err)
	}
}

// Save model to disk
func (m *SimModel) Save() {
	if m.changed {
		fmt.Println("Saving simulation properties...")
		indentedJSON, _ := json.MarshalIndent(m.Sim, "", "  ")

		dataPath, err := filepath.Abs(m.relativePath)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(dataPath+m.file, indentedJSON, 0644)
		if err != nil {
			log.Fatalln("ERROR:", err)
		}

		m.Clean()
		fmt.Println("Simulation properties saved")
	}
}
