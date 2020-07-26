package streams

import (
	"math"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
)

// Stimulus streams are typically streams sourced from a file
// during developement. In practice though, stimulus comes from
// actual stimulus sources such as images, INs or other neurons.
//
// When patterns are emitted there is gap (aka interval) between patterns.
// The Inter-Pattern-Interval (IPI) can form several ways:
// 1) Randomly with a minimum interval size.
// 2) Regularly based on a frequency (Hz). IPI is implied via the frequency.
// 3) Poisson distributed IPI.

type stimulusStream struct {
	pattern []int

	// Inter-Pattern-Interval (IPI)
	ipi int

	// How often the pattern in presented (in Hertz)
	frequency int

	// Sub duration count
	count int

	presentingPattern bool
	bitIdx            int
}

// Example Format:
// ....|.     <-- A pattern is just a single row
// ...|..
// |..|..
// .|....
// ....|.
// ....|.
// |.....
// .....|
// ..|...
// .|....

// NewStimulusStream ...
func NewStimulusStream(pattern []int, frequency int) api.IBitStream {
	o := new(stimulusStream)
	o.pattern = pattern
	o.frequency = frequency

	patternLength := len(pattern)
	// frequency = patterns/second or pattern/1000ms
	milliseconds := 1000.0 // convert to milliseconds
	period := 1.0 / float64(frequency)
	o.ipi = int(math.Round(period*milliseconds)) - patternLength

	// fmt.Println("-------------------------------")
	// fmt.Println("StimulusStream properties:")
	// fmt.Println("period: ", period)
	// fmt.Println("pattern presented every: ", period*1000, " ms")
	// fmt.Println("patternLength: ", patternLength)
	// fmt.Println("ipi: ", o.ipi)
	// fmt.Println("-------------------------------")

	o.Reset()

	return o
}

// Reset ...
func (s *stimulusStream) Reset() {
	s.count = 0
	s.presentingPattern = true
	s.bitIdx = 0
}

// Step ...
// frequency is specified in Hz, for example if Hz = 10 then the pattern
// is presented every 1/10 of a second or every 100ms. If the TimeScale
// is 100us then presentation can be thought of as 10000us.
// The time layout is as follows:
// |---------- 1 presentation ---------|---------- 2 presentation ---------|...
// |----- Pattern -----|----- IPI -----|----- Pattern -----|----- IPI -----|...
//
// If the frequency is 10Hz (period 100ms) and the pattern length is 30ms
// then cycle layout is as follows:
// |30ms pattern|70ms IPI|30ms pattern|70ms IPI|30ms pattern|70ms IPI|...
//
// step should be called only once for the pattern and NOT for each synapse.
func (s *stimulusStream) Step() {
	s.count--

	if s.presentingPattern {
		if s.count <= 0 {
			// Reset counter to IPI
			s.count = s.ipi
			s.presentingPattern = false
		} else {
			s.bitIdx++
		}
	} else {
		if s.count <= 0 {
			// Reset counter to Pattern
			s.count = len(s.pattern)
			// Reset pattern for next presentation
			s.bitIdx = 0
			s.presentingPattern = true
		}
	}
}

// Output ...
func (s *stimulusStream) Output() int {
	if s.presentingPattern {
		return s.pattern[s.bitIdx]
	}

	return 0
}

// Update changes the stream's properties
func (s *stimulusStream) Update(mod api.IModel) {
	// conData, _ := mod.Data().(*model.ConfigJSON)

}
