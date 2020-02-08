package stimulus

import (
	"fmt"
	"image"

	// You must register PNG otherwise you get "unknown format" err during decoding.
	"image/color"
	_ "image/png" // Register PNG format
	"os"

	"github.com/wdevore/Deuron8-Go/interfaces"
)

// There are several possible types of stimulus for an AI system, however,
// Deuron8 focuses on visual images.
// The stimulus is a square image that is scanned line by line.
// Each pixel on a line is transformed into a row of spikes that
// impact on several Synapses on different Dendrites on different Somas.
//
// pixel  -->.
//      | ||  | |||   <-- a scanline
//        ||| |   |
//       |   ||  ||
//
// Each spike needs to be routed to a synapse on a Cell. Each Cell
// is a random distance from the stimulus. The distance is represented as
// a delay. The delay is contained in source stimulus.

// API is the runtime configuration
var API interfaces.IVisStimInput

// Implements interface
type visualStimulus struct {
	// The image sourcing the stimulus
	image image.Image
	// Intensity as grayscale image
	gray *image.Gray

	expandEnabled bool
	expand        string
	size          string

	// Stimulus at a pixel
	stimulus []int
}

// New constructs an IVisStimInput object
func New() interfaces.IVisStimInput {
	o := new(visualStimulus)
	o.expandEnabled = false
	return o
}

func (vs *visualStimulus) Configure(imageFile string) error {
	infile, err := os.Open(imageFile)
	if err != nil {
		return err
	}
	defer infile.Close()

	// Decode will figure out what type of image is in the file on its own.
	// We just have to be sure all the image packages we want are imported.
	vs.image, _, err = image.Decode(infile)
	if err != nil {
		return err
	}

	// Convert image to grayscale
	bounds := vs.image.Bounds()
	vs.gray = image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			vs.gray.Set(x, y, vs.image.At(x, y))
		}
	}

	return nil
}

func (vs *visualStimulus) SetSize(size string) {
	vs.size = size
}

func (vs *visualStimulus) EnableExpand(enable bool) {
	vs.expandEnabled = enable
}

func (vs *visualStimulus) SetExpand(expand string) {
	vs.expand = expand
}

func (vs *visualStimulus) SetStimulusAt(x, y int) {
	// Using x, y we can find Color and intensity components
	c := vs.image.At(x, y).(color.NRGBA) // A Color tuple
	i := vs.gray.At(x, y).(color.Gray)   // Intensity

	vs.SetSize("10")
	exp := vs.Expand(x)
	exp = exp + vs.Expand(y)
	vs.SetSize("8")
	exp = exp + vs.Expand(int(c.R))
	exp = exp + vs.Expand(int(c.G))
	exp = exp + vs.Expand(int(c.B))
	exp = exp + vs.Expand(int(i.Y))

	vs.stimulus = make([]int, len(exp))
	for i, b := range exp {
		if b == '1' {
			vs.stimulus[i] = 1
		} else {
			vs.stimulus[i] = 0
		}
	}
}

func (vs *visualStimulus) GetStimulus() []int {
	return vs.stimulus
}

func (vs *visualStimulus) GetStimulusComp(value int) []int {
	exp := vs.Expand(value)

	spk := make([]int, len(exp))
	for i, b := range exp {
		if b == '1' {
			spk[i] = 1
		} else {
			spk[i] = 0
		}
	}

	return spk
}

// Expand expands a value, for example: ("00")
func (vs *visualStimulus) Expand(value int) string {
	var bin string
	bin = fmt.Sprintf("%0"+vs.size+"b", value)

	if vs.expandEnabled {
		exp := ""

		for _, c := range bin {
			p := string(c) + vs.expand
			exp += p
		}

		// fmt.Println(bin)
		// fmt.Println(exp)
		// fmt.Println(len(exp))

		return exp
	} else {
		return bin
	}
}
