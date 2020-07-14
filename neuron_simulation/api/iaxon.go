package api

// IAxon routes spikes to synapses
type IAxon interface {
	Reset()
	Input(int)
	Output() int
	Step()

	Set(int)
	SetNoDelay()
	SetToMaxDelay()
	SetToHalfDelay()
}
