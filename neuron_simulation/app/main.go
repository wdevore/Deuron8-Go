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
	// Construct environment
	// -------------------------------------------------------------
	environment := NewEnvironment("../../", "/neuron_simulation/data/")

	moData, _ := environment.Config().Data().(*model.ConfigJSON)

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
	run(platform, renderer, environment)
}
