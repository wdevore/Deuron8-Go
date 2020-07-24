package main

import (
	"fmt"
	"time"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/app/gui"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/cellsimulation"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/graphs"
)

// Platform covers mouse/keyboard/gamepad inputs, cursor shape, timing, windowing.
type Platform interface {
	// ShouldStop is regularly called as the abort condition for the program loop.
	ShouldStop() bool
	// ProcessEvents is called once per render loop to dispatch any pending events.
	ProcessEvents()
	// DisplaySize returns the dimension of the display.
	DisplaySize() [2]float32
	// FramebufferSize returns the dimension of the framebuffer.
	FramebufferSize() [2]float32
	// NewFrame marks the begin of a render pass. It must update the imgui IO state according to user input (mouse, keyboard, ...)
	NewFrame()
	// PostRender marks the completion of one render pass. Typically this causes the display buffer to be swapped.
	PostRender()
	// ClipboardText returns the current text of the clipboard, if available.
	ClipboardText() (string, error)
	// SetClipboardText sets the text as the current text of the clipboard.
	SetClipboardText(text string)
}

type clipboard struct {
	platform Platform
}

func (board clipboard) Text() (string, error) {
	return board.platform.ClipboardText()
}

func (board clipboard) SetText(text string) {
	board.platform.SetClipboardText(text)
}

// Renderer covers rendering imgui draw data.
type Renderer interface {
	// PreRender causes the display buffer to be prepared for new output.
	PreRender(clearColor [3]float32)
	// Render draws the provided imgui draw data.
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.DrawData)
}

const (
	millisPerSecond = 1000
	sleepDuration   = time.Millisecond * 33
)

// run implements the main program loop of the app. It returns when the platform signals to stop.
func run(p Platform, r Renderer, environment api.IEnvironment) {
	imgui.CurrentIO().SetClipboard(clipboard{platform: p})

	clearColor := [3]float32{0.25, 0.25, 0.25}

	spikeGraph := graphs.NewSpikeGraph()
	surgeGraph := graphs.NewSynapseSurgeGraph()
	pspGraph := graphs.NewSynapsePspGraph()

	ch := make(chan string)

	simulator := cellsimulation.NewSimulator(environment)
	simulator.Build()

	// Start simulation thread. It will idle by default.
	go simulator.Run(ch)

	simThreadKilled := false

	// -------------------------------------------------------------
	// Now start main GUI loop
	// -------------------------------------------------------------
	for !p.ShouldStop() {
		p.ProcessEvents()
		p.NewFrame()

		imgui.NewFrame()

		// ---------------------------------------------------------
		// Draw Graphs
		// ---------------------------------------------------------
		vertPos := 40
		spikeGraph.Draw(environment, vertPos)

		vertPos += graphs.SpikePanelHeight + 20
		surgeGraph.Draw(environment, vertPos)

		vertPos += graphs.SurgePanelHeight
		pspGraph.Draw(environment, vertPos)
		// ---------------------------------------------------------

		gui.BuildGui(environment)

		imgui.EndFrame()
		// ---------------------------------------------------------

		// Rendering
		imgui.Render() // This call only creates the draw data list. Actual rendering to framebuffer is done below.

		r.PreRender(clearColor)

		// A this point, the application could perform its own rendering...
		// app.RenderScene()

		r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.RenderedDrawData())

		p.PostRender()

		if !simThreadKilled {
			if environment.IsCmdIssued() {
				// fmt.Println("Issuing Cmd: ", environment.Cmd())
				ch <- environment.Cmd() // sends message to channel
				if environment.Cmd() == "killSim" {
					fmt.Println("Killing simulator thread")
					simThreadKilled = true
				}
				environment.CmdIssued()
			}
		} else {
			if environment.IsCmdIssued() {
				fmt.Println("WARNING!! Simulation thread was killed! No commands accepted.")
			}
		}

		// sleep to avoid 100% CPU usage
		<-time.After(sleepDuration)
	}

	fmt.Println("Exiting application")
	gui.Shutdown(environment)
}
