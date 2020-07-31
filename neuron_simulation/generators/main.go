package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/misc"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

func main() {

	// Read config to determine what needs to be done
	config := NewConfigModel("config.json")

	c := config.Data().(*ConfigJSON)
	d := c.Presets

	mod := model.SynapsesJSON{}

	for i := 0; i < c.SynapseCount; i++ {
		syn := model.SynapseJSON{}

		syn.ID = i
		syn.TaoP = calcV(&d.TaoP)
		syn.TaoN = calcV(&d.TaoN)
		syn.Mu = calcV(&d.Mu)
		syn.Distance = calcV(&d.Distance)
		syn.Lambda = calcV(&d.Lambda)
		syn.Amb = calcV(&d.Amb)
		syn.W = calcV(&d.W)
		syn.Alpha = calcV(&d.Alpha)
		syn.LearningRateSlow = calcV(&d.LearningRateSlow)
		syn.LearningRateFast = calcV(&d.LearningRateFast)
		syn.TaoI = calcV(&d.TaoI)
		syn.Ama = calcV(&d.Ama)

		mod.Synapses = append(mod.Synapses, &syn)
	}

	indentedJSON, _ := json.MarshalIndent(mod, "", "  ")

	dataPath, err := filepath.Abs(c.TargetPath)
	if err != nil {
		panic(err)
	}

	file := dataPath + c.OutputFile
	fmt.Println("Saving synapses presets to " + file)

	err = ioutil.WriteFile(file, indentedJSON, 0644)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	fmt.Println("Presets saved")

}

func calcV(preset *Preset) float64 {
	if preset.Enabled {
		l := misc.Linear(preset.Min, preset.Max, preset.Center)
		r := rand.Float64()

		if r > l {
			// Center -> Max wins
			return misc.Lerp(preset.Center, preset.Max, rand.Float64())
		}

		// Min -> Center wins
		return misc.Lerp(preset.Min, preset.Center, rand.Float64())
	}

	return preset.Center
}
