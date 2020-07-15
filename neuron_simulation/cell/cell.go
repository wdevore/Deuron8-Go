package cell

import (
	"github.com/wdevore/Deuron8-Go/neuron_simulation/api"
	"github.com/wdevore/Deuron8-Go/neuron_simulation/model"
)

// Cell is an encapsulation of a neuron (soma, Cell etc.)
// The Cell is the implementation of a Neuron.
// The cell generates streams of spikes and thus is a bit stream too.
// The stream is written to disc in spans.
type Cell struct {
	soma api.ISoma

	simJ     *model.SimJSON
	simModel api.IModel
}

// NewCell creates a new Cell
func NewCell(simModel api.IModel, soma api.ISoma) api.ICell {
	o := new(Cell)
	o.simModel = simModel

	o.soma = soma

	simJ, ok := simModel.Data().(*model.SimJSON)

	if !ok {
		panic("Cell: can't cast simModel to SimJSON")
	}

	o.simJ = simJ

	return o
}

// Initialize Cell
func (c *Cell) Initialize() {
	c.soma.Initialize()
}

// Reset Cell
func (c *Cell) Reset() {
	c.soma.Reset()
}

// Integrate is the actual integration
func (c *Cell) Integrate(spanT, t int) (spike int) {
	return c.soma.Integrate(spanT, t)
}

// Output of Cell
func (c *Cell) Output() int {
	return c.soma.Output()
}
