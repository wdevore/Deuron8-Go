package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/wdevore/Deuron8-Go/interfaces"
)

// API is the runtime configuration
var API interfaces.IConfig

// The JSON data structure
type configJSON struct {
	ErrLog    string `json:"ErrLog"`
	InfoLog   string `json:"InfoLog"`
	ExitState string `json:"ExitState"`
	LogRoot   string `json:"LogRoot"`
}

type configuration struct {
	dirty      bool
	conf       configJSON
	path       string
	configFile string
}

// New construct an IConfig object
func New(configFile string) interfaces.IConfig {
	o := new(configuration)
	o.dirty = false
	o.configFile = configFile

	// dir, err := filepath.Abs(filepath.Dir(""))
	confPath, err := filepath.Abs("")
	if err != nil {
		log.Fatal(err)
	}

	o.path = confPath

	jsonFile, err := os.Open(confPath + "/" + configFile)
	if err != nil {
		log.Fatalln("ERROR:", err)
		return nil
	}

	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&o.conf)

	if err != nil {
		log.Fatalln("ERROR:", err)
		return nil
	}

	return o
}

// Save persists the current config to json file.
func (c *configuration) Save() {
	if !c.dirty {
		fmt.Print("nothing to save")
		return
	}

	indentedJSON, _ := json.MarshalIndent(c.conf, "", "  ")
	err := ioutil.WriteFile(c.path+c.configFile, indentedJSON, 0644)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
}

// ErrLogFileName is the name of the error log file.
func (c *configuration) ErrLogFileName() string {
	return c.conf.ErrLog
}

// InfoLogFileName is the name of the info log file.
func (c *configuration) InfoLogFileName() string {
	return c.conf.InfoLog
}

// LogRoot is the base path to where log files are located.
func (c *configuration) LogRoot() string {
	return c.conf.LogRoot
}

// ExitState indicates what the last state the
// simulation was in when deuron exited.
// Values:
//   Terminated = user quit simulation while it was inprogress
//   Completed = sim terminated on its own
//   Crashed = sim died
//   Paused = user paused simulation and exited
//   Exited = user exited when no simulation was running
func (c *configuration) ExitState() string {
	return c.conf.ExitState
}

// SetExitState sets a value upon deuron exit.
func (c *configuration) SetExitState(state string) {
	c.conf.ExitState = state
	c.dirty = true
}
