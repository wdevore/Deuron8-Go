package api

// IEnvironment is the runtime environment
type IEnvironment interface {
	Config() IModel
	Sim() IModel
	Synapses() IModel

	Samples() ISamples

	Stimulus() [][]int
	StimulusAt(int) []int
	StimulusCount() int

	IssueCmd(string)
	IsCmdIssued() bool
	CmdIssued()
	Cmd() string

	// Because the simulator runs in a coroutine the GUI
	// can't attempt to access any samples until the simulation is complete.
	IsRunning() bool
	Run(bool)
}
