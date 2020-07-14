package main

import (
	"fmt"
	"os"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/app/platforms"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/app/renderers"
)

func main() {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := platforms.NewGLFW("Deuron8-Go", io, platforms.GLFWClientAPIOpenGL3)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderers.NewOpenGL3(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	Run(platform, renderer)
}
