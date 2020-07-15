package main

import (
	"fmt"
	"os"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/app/platforms"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/app/renderers"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

func main() {
	fmt.Println("Welcome to Deuron8 Go edition")

	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	// -------------------------------------------------------------
	// Load application property settings
	// -------------------------------------------------------------
	config := model.NewConfigModel("../../", "/neuron_simulation/jsondata/config.json")

	moData, ok := config.Data().(*model.ConfigJSON)

	if ok {
		fmt.Println("Config loaded.")
		fmt.Println("AutoSave: ", moData.AutoSave)
	}

	sim := model.NewSimModel("../../", "/neuron_simulation/jsondata/sim_model.json")

	simData, ok := sim.Data().(*model.SimJSON)

	if ok {
		fmt.Println("Default sim_model loaded.")
		fmt.Println("Synapses: ", simData.Synapses)
	}

	// -------------------------------------------------------------
	// GLFW Window
	// -------------------------------------------------------------
	platform, err := platforms.NewGLFW("Deuron8-Go", moData.WindowWidth, moData.WindowHeight, io, platforms.GLFWClientAPIOpenGL3)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	// -------------------------------------------------------------
	// Renderer used by GLFW
	// -------------------------------------------------------------
	renderer, err := renderers.NewOpenGL3(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	// Finally run main gui application
	Run(platform, renderer, config, sim)
}
