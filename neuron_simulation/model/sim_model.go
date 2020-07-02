package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

// SimModel is the simulation presistance data
type SimModel struct {
	Sim SimJSON
}

// NewSimModel creates and loads sim data
func NewSimModel(relativePath, file string) api.IModel {
	o := new(SimModel)

	o.Sim = SimJSON{}

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

	err = json.Unmarshal(bytes, &o.Sim)
	if err != nil {
		panic(err)
	}

	return o
}

// Data returns the json loaded app data
func (m *SimModel) Data() api.IModelData {
	return &m.Sim
}
