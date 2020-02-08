package stimulustests

// go test -v stimulus_test.go

import (
	"fmt"
	"testing"

	"github.com/wdevore/Deuron8-Go/config"
	logg "github.com/wdevore/Deuron8-Go/log"
	"github.com/wdevore/Deuron8-Go/simulation/stimulus"
)

func init() {
	config.API = config.New("../config/config.json")
	logg.API = logg.New(config.API)
}
func TestMain(t *testing.T) {
	defer config.API.Save()
	defer logg.API.Close()

	logg.API.LogInfo("Starting stim test")

	runStimTest(t)
}

func runStimTest(t *testing.T) {
	vis := stimulus.New()
	err := vis.Configure("../simulation/stimulus/letter_A.png")
	if err != nil {
		logg.API.LogError(err.Error())
		return
	}

	vis.SetExpand("00")
	vis.SetSize("8")
	// 00000010 00001001 00000100
	x := vis.Expand(45)
	logg.API.LogInfo(x)

	// 100100100100100000100000000000
	vis.SetSize("10")
	x = vis.Expand(1000)
	logg.API.LogInfo(x)

	// 1 0 0 1 0 0 1 0 0 1 0 0 1 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0
	xa := vis.GetStimulusComp(1000)
	fmt.Println(xa)

	vis.SetStimulusAt(454, 500)
	stim := vis.GetStimulus()
	fmt.Println(stim)
	fmt.Println(len(stim))
}
