package plotstests

import (
	"image/color"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func TestMain(t *testing.T) {
	plotPoisson(t)
}

func plotPoisson(t *testing.T) {
	// lambda = k (events/time * time-interval)     / n (units)
	// n = k / l = 50
	poi := distuv.Poisson{Lambda: 4.0}
	// Draw some random values from the standard normal distribution.
	v := make(plotter.Values, 1000)
	for i := range v {
		v[i] = poi.Rand()
	}

	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Poisson Histogram"

	// Create a histogram of our values drawn
	// from the random function.
	h, err := plotter.NewHist(v, 16)
	if err != nil {
		panic(err)
	}
	// Normalize the area under the histogram to
	// sum to one.
	h.Normalize(1) // Vertical axis will be from 0->1
	p.Add(h)

	// The poisson distribution function
	// norm := plotter.NewFunction(distuv.Poisson{Lambda: 1.0}.Prob)
	norm := plotter.NewFunction(poi.Prob)
	norm.Color = color.RGBA{G: 255, A: 255}
	norm.Width = vg.Points(1)
	p.Add(norm)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "poisson.png"); err != nil {
		panic(err)
	}
}

func plotHistogram(t *testing.T) {
	// Draw some random values from the standard
	// normal distribution.
	rand.Seed(int64(0))
	v := make(plotter.Values, 10000)
	for i := range v {
		v[i] = rand.NormFloat64()
	}

	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Histogram"

	// Create a histogram of our values drawn
	// from the standard normal.
	h, err := plotter.NewHist(v, 16)
	if err != nil {
		panic(err)
	}
	// Normalize the area under the histogram to
	// sum to one.
	h.Normalize(1)
	p.Add(h)

	// The normal distribution function
	norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	p.Add(norm)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "hist.png"); err != nil {
		panic(err)
	}
}

func plotRandom(t *testing.T) {
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"First", randomPoints(15),
		"Second", randomPoints(15),
		"Third", randomPoints(15))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
