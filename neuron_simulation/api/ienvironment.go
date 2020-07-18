package api

// IEnvironment is the runtime environment
type IEnvironment interface {
	Config() IModel
	Sim() IModel
	Samples() ISamples
	Stimulus() [][]int
	StimulusAt(int) []int

	IssueCmd(string)
	IsCmdIssued() bool
	CmdIssued()
	Cmd() string

	AutoStop(bool)
	IsAutoStop() bool
}
