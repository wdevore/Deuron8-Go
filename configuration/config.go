package configuration

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/wdevore/Deuron8-Go/interfaces"
)

// Config is a JSON decoded object. See config.json
type Config struct {
	ErrLog string `json:"ErrLog"`
}

// Configuration is the main configuration of the app.
type Configuration struct {
	config Config
}

// New construct an IConfig object
func New() interfaces.IConfig {
	o := new(Configuration)

	const configFile = "./config.json"
	confPath, err := filepath.Abs("configuration")
	// dir, err := filepath.Abs(filepath.Dir(""))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)

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

// ErrLogFileName is the name of the error log file.
func (c *Configuration) ErrLogFileName() string {
	return c.config.ErrLog
}
