package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

// Model is the main presistance data
type Model struct {
	Config ConfigJSON

	relativePath string
}

// NewModel creates and loads app data
func NewModel(relativePath, file string) api.IModel {
	o := new(Model)

	o.Config = ConfigJSON{}

	o.relativePath = relativePath

	dataPath, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}

	eConfFile, err := os.Open(dataPath + "/neuron_simulation/" + file)
	if err != nil {
		panic(err)
	}

	defer eConfFile.Close()

	bytes, err := ioutil.ReadAll(eConfFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, o)
	if err != nil {
		panic(err)
	}

	return o
}

// Data returns the json loaded app data
func (m *Model) Data() api.IModelData {
	return &m.Config
}
