package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Model is the main presistance data
type Model struct {
	Config ConfigJSON

	file string
}

// NewConfigModel creates and loads app data
func NewConfigModel(file string) *Model {
	o := new(Model)

	o.Config = ConfigJSON{}

	o.file = file

	o.Load()

	return o
}

// Data returns the json loaded app data
func (m *Model) Data() interface{} {
	return &m.Config
}

// Load model from disk
func (m *Model) Load() {
	dataPath, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	eConfFile, err := os.Open(dataPath + "/" + m.file)
	if err != nil {
		panic(err)
	}

	defer eConfFile.Close()

	bytes, err := ioutil.ReadAll(eConfFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &m.Config)
	if err != nil {
		panic(err)
	}
}
