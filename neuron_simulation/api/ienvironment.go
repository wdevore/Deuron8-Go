package api

// IEnvironment is the runtime environment
type IEnvironment interface {
	Config() IModel
	Sim() IModel
	Samples() ISamples
	Stimulus() [][]int
	StimulusAt(int) []int
}
