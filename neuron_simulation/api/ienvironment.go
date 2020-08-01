package api

const (
	WeightBoundingHard = 0
	WeightBoundingSoft = 1
)

// IEnvironment is the runtime environment
type IEnvironment interface {
	Config() IModel
	Sim() IModel
	SynapticModel() IModel

	AddSynapse(ISynapse)
	Synapses() []ISynapse

	Samples() ISamples

	Stimulus() [][]int
	StimulusAt(int) []int
	StimulusCount() int

	IssueCmd(string)
	IsCmdIssued() bool
	CmdIssued()
	Cmd() string
	Parms() string
	SetParms(parms string)

	// Because the simulator runs in a coroutine the GUI
	// can't attempt to access any samples until the simulation is complete.
	IsRunning() bool
	Run(bool)

	// Runtime properties
	RandomizerField() int
	SetRandomizerField(int)

	MinimumRangeValue() float64
	SetMinimumRangeValue(float64)

	MaximumRangeValue() float64
	SetMaximumRangeValue(float64)

	CenterRangeValue() float64
	SetCenterRangeValue(float64)

	InitialWeightValues() int
	SetInitialWeightValues(int)

	WeightBounding() int
	SetWeightBounding(int)
}
