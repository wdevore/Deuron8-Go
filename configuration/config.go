package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/wdevore/Deuron8-Go/interfaces"
)

// Config is a JSON decoded object. See config.json
type Config struct {
	ErrLog    string `json:"ErrLog"`
	InfoLog   string `json:"InfoLog"`
	ExitState string `json:"ExitState"`
	LogRoot   string `json:"LogRoot"`
}

// Configuration is the main configuration of the app.
type Configuration struct {
	config Config
	path   string
}

const configFile = "/config.json"

// New construct an IConfig object
func New() interfaces.IConfig {
	o := new(Configuration)

	// dir, err := filepath.Abs(filepath.Dir(""))
	confPath, err := filepath.Abs("configuration")
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

	err = json.NewDecoder(jsonFile).Decode(&o.config)

	if err != nil {
		log.Fatalln("ERROR:", err)
		return nil
	}

	return o
}

// Save persists the current config to json file.
func (c *Configuration) Save() {
	indentedJSON, _ := json.MarshalIndent(c.config, "", "  ")
	err := ioutil.WriteFile(c.path+configFile, indentedJSON, 0644)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
}

// ErrLogFileName is the name of the error log file.
func (c *Configuration) ErrLogFileName() string {
	return c.config.ErrLog
}

// InfoLogFileName is the name of the info log file.
func (c *Configuration) InfoLogFileName() string {
	return c.config.ErrLog
}

// LogRoot is the base path to where log files are located.
func (c *Configuration) LogRoot() string {
	return c.config.LogRoot
}

// ExitState indicates what the last state the
// simulation was in when deuron exited.
// Values:
//   Terminated = user quit simulation while it was inprogress
//   Completed = sim terminated on its own
//   Crashed = sim died
//   Paused = user paused simulation
func (c *Configuration) ExitState() string {
	return c.config.ExitState
}

// SetExitState sets a value upon deuron exit.
func (c *Configuration) SetExitState(state string) {
	c.config.ExitState = state
}
