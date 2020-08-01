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
		_, syn.TaoP = calcV(&d.TaoP)
		_, syn.TaoN = calcV(&d.TaoN)
		_, syn.Mu = calcV(&d.Mu)
		_, syn.Distance = calcV(&d.Distance)
		_, syn.Lambda = calcV(&d.Lambda)
		_, syn.Amb = calcV(&d.Amb)
		syn.Excititory, syn.W = calcV(&d.W)
		_, syn.Alpha = calcV(&d.Alpha)
		_, syn.LearningRateSlow = calcV(&d.LearningRateSlow)
		_, syn.LearningRateFast = calcV(&d.LearningRateFast)
		_, syn.TaoI = calcV(&d.TaoI)
		_, syn.Ama = calcV(&d.Ama)

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

func calcV(preset *Preset) (bool, float64) {
	if preset.Enabled {
		l := misc.Linear(preset.Min, preset.Max, preset.Center)
		r := rand.Float64()

		if r > l {
			// Center -> Max wins
			return true, misc.Lerp(preset.Center, preset.Max, rand.Float64())
		}

		// Min -> Center wins
		return false, misc.Lerp(preset.Min, preset.Center, rand.Float64())
	}

	return true, preset.Center
}
