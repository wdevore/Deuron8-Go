package streams

import (
	"math"

	"golang.org/x/exp/rand"

	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"gonum.org/v1/gonum/stat/distuv"
)

// https://support.minitab.com/en-us/minitab-express/1/help-and-how-to/basic-statistics/probability-distributions/supporting-topics/distributions/poisson-distribution/
// Poisson is commonly used for modelling the number of occurrences
// of an event within a particular time interval.
// For example, we may want an average of 20 spikes to occur within
// a 1 sec interval.
// Note: where the spikes occur within the interval is random, but
// we expect to average 20 spikes within that interval.

// What is rate of occurrence?
// The rate of occurrence equals the mean (λ) divided by the dimension
// of your observation space (interval). It is useful for comparing Poisson
// counts collected in different observation spaces.
// For example, Switchboard A receives 50 telephone calls in 5 hours,
// and Switchboard B receives 80 calls in 10 hours.
// You cannot directly compare these values because their observation
// spaces are different.
// You must calculate the occurrence rate to compare these counts.
// The rate for Switchboard A is (50 calls / 5 hours) = 10 calls/hour.
// The rate for Switchboard B is (80 calls / 10 hours) = 8 calls/hour.

// Generating:
// If you have a Poisson process with rate parameter
// L (meaning that, long term, there are L arrivals per second),
// then the inter-arrival times are exponentially distributed with
// mean 1/L.
// So the PDF is f(t) = -L*exp(-Lt),
// and the CDF is F(t) = Prob(T < t) = 1 - exp(-Lt).
// So your problem changes to: how do I generate a random number t
// with distribution F(t) = 1 - \exp(-Lt)?

// Assuming the language you are using has a function (let's call it rand())
// to generate random numbers uniformly distributed between 0 and 1,
// the inverse CDF technique reduces to calculating: -log(rand()) / L

type poissonStream struct {
	poisson distuv.Poisson

	// The Interspike interval (ISI) is a counter
	// When the counter reaches 0 a spike is placed on the output
	// for single pass.
	isi int64

	// λ is the shape parameter which indicates the 'average' number of
	// events in the given time interval
	averagePerInterval float64

	seed uint64
}

// NewPoissonStream creates a new poisson distributed stream of spikes
func NewPoissonStream(seed uint64, averagePerInterval float64) api.IBitStream {
	o := new(poissonStream)

	o.seed = seed
	o.averagePerInterval = averagePerInterval

	o.Reset()
	return o
}

// Reset ...
func (p *poissonStream) Reset() {
	psource := rand.NewSource(uint64(p.seed))
	p.poisson = distuv.Poisson{Lambda: p.averagePerInterval, Src: psource}
	p.isi = p.next()
}

// Step ...
func (p *poissonStream) Step() {
}

// Output ...
func (p *poissonStream) Output() int {
	if p.isi == 0 {
		p.isi = p.next()
		return 1
	}

	p.isi = -1
	return 0
}

// Create an event per interval of time, for example, spikes in a 1 sec span.
// A firing rate given in rate/ms, for example, 0.2 in 1ms (0.2/1)
// or 200 in 1sec (200/1000ms)
func (p *poissonStream) next() int64 {
	isiF := -math.Log(1.0-p.poisson.Rand()) / p.averagePerInterval
	return int64(math.Round(isiF))
}
