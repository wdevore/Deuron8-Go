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

// Model is the main presistance data
type Model struct {
	Config ConfigJSON

	relativePath string
	file         string

	samples api.ISamples

	changed bool
}

// NewConfigModel creates and loads app data
func NewConfigModel(relativePath, file string) api.IModel {
	o := new(Model)

	o.Config = ConfigJSON{}

	o.relativePath = relativePath
	o.file = file

	dataPath, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}

	eConfFile, err := os.Open(dataPath + file)
	if err != nil {
		panic(err)
	}

	defer eConfFile.Close()

	bytes, err := ioutil.ReadAll(eConfFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &o.Config)
	if err != nil {
		panic(err)
	}

	// o.samples =

	return o
}

// Data returns the json loaded app data
func (m *Model) Data() interface{} {
	return &m.Config
}

// SetActiveSynapse not used
func (m *Model) SetActiveSynapse(id int) {
}

// Samples returns the simulation samples
func (m *Model) Samples() api.ISamples {
	return m.samples
}

// Changed marks model dirty
func (m *Model) Changed() {
	m.changed = true
}

// Clean marks model NOT-dirty
func (m *Model) Clean() {
	m.changed = false
}

// Save model to disk
func (m *Model) Save() {
	if m.changed {
		fmt.Println("Saving application properties...")
		indentedJSON, _ := json.MarshalIndent(m.Config, "", "  ")

		dataPath, err := filepath.Abs(m.relativePath)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(dataPath+m.file, indentedJSON, 0644)
		if err != nil {
			log.Fatalln("ERROR:", err)
		}

		m.Clean()
		fmt.Println("Application properties saved")
	}
}