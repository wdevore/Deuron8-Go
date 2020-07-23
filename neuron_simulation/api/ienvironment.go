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
}
