package stimulus

import (
	"fmt"
	"image"

	// You must register PNG otherwise you get "unknown format" err during decoding.
	"image/color"
	_ "image/png" // Register PNG format
	"os"

	"github.com/wdevore/Deuron8-Go/api"
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
var API api.IVisStimInput

const cacheSize = 1024

// Implements interface
type visualStimulus struct {
	// The image sourcing the stimulus
	image image.Image
	// Intensity as grayscale image
	gray *image.Gray

	expandEnabled bool
	expand        string
	size          string
	expandLen     int

	// List of preconverted bit-strings into int slices
	bits            [][]int
	preCacheEnabled bool

	// Stimulus at a pixel
	stimulus []int
}

// New constructs an IVisStimInput object
func New() api.IVisStimInput {
	o := new(visualStimulus)
	o.expandEnabled = false
	o.preCacheEnabled = false
	o.computeLen()
	o.stimulus = make([]int, o.expandLen)
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

	vs.SetSize("8")
	vs.bits = make([][]int, cacheSize)

	// Preconvert bit strings
	for i := 0; i < 255; i++ {
		b := vs.GetStimulusComp(i)
		vs.bits[i] = []int{b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]}
	}

	vs.SetSize("10")
	for i := 256; i < cacheSize; i++ {
		b := vs.GetStimulusComp(i)
		vs.bits[i] = []int{b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9]}
	}

	vs.preCacheEnabled = true

	return nil
}

func (vs *visualStimulus) SetSize(size string) {
	vs.size = size
}

func (vs *visualStimulus) EnableExpand(enable bool) {
	vs.expandEnabled = enable
}

func (vs *visualStimulus) EnablePreCache(enable bool) {
	vs.preCacheEnabled = enable
}

func (vs *visualStimulus) SetExpand(expand string) {
	vs.expand = expand
}

func (vs *visualStimulus) computeLen() {
	vs.SetSize("10")
	exp := vs.Expand(0)
	exp = exp + vs.Expand(0)
	vs.SetSize("8")
	exp = exp + vs.Expand(0)
	exp = exp + vs.Expand(0)
	exp = exp + vs.Expand(0)
	exp = exp + vs.Expand(0)

	vs.expandLen = len(exp)
}

// Each bit feeds into a synapse on a dendrite within a neaboring
// region.
func (vs *visualStimulus) SetStimulusFast(x, y int) {
	// Using x, y we can find Color and intensity components
	c := vs.image.At(x, y).(color.NRGBA) // A Color tuple
	i := vs.gray.At(x, y).(color.Gray)   // Intensity

	s := vs.stimulus

	co := vs.GetStimulusComp(x) // 10bits
	s[0] = co[0]
	s[1] = co[1]
	s[2] = co[2]
	s[3] = co[3]
	s[4] = co[4]
	s[5] = co[5]
	s[6] = co[6]
	s[7] = co[7]
	s[8] = co[8]
	s[9] = co[9]

	co = vs.GetStimulusComp(y) // 10bits
	s[10] = co[0]
	s[11] = co[1]
	s[12] = co[2]
	s[13] = co[3]
	s[14] = co[4]
	s[15] = co[5]
	s[16] = co[6]
	s[17] = co[7]
	s[18] = co[8]
	s[19] = co[9]

	co = vs.GetStimulusComp(int(c.R)) // 8bits
	s[20] = co[0]
	s[21] = co[1]
	s[22] = co[2]
	s[23] = co[3]
	s[24] = co[4]
	s[25] = co[5]
	s[26] = co[6]
	s[27] = co[7]

	co = vs.GetStimulusComp(int(c.G))
	s[28] = co[0]
	s[29] = co[1]
	s[30] = co[2]
	s[31] = co[3]
	s[32] = co[4]
	s[33] = co[5]
	s[34] = co[6]
	s[35] = co[7]

	co = vs.GetStimulusComp(int(c.B))
	s[36] = co[0]
	s[37] = co[1]
	s[38] = co[2]
	s[39] = co[3]
	s[40] = co[4]
	s[41] = co[5]
	s[42] = co[6]
	s[43] = co[7]

	// Intensity is has only 3 patterns: Low, med and high
	co = vs.GetStimulusComp(int(i.Y)) // 8bits
	s[44] = co[0]
	s[45] = co[1]
	s[46] = co[2]
	s[47] = co[3]
	s[48] = co[4]
	s[49] = co[5]
	s[50] = co[6]
	s[51] = co[7]
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
	if vs.preCacheEnabled {
		return vs.bits[value]
	}

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

		fmt.Println(bin)
		fmt.Println(exp)
		fmt.Println(len(exp))

		return exp
	}

	return bin
}
