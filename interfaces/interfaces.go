package interfaces

// ILogger is a general logging interface
type ILogger interface {
	LogError(string)
	LogInfo(string)
	Close()
}

// IConfig holds configuration and runtime properties.
type IConfig interface {
	ErrLogFileName() string
	InfoLogFileName() string
	LogRoot() string

	ExitState() string
	SetExitState(string)

	Save()
}

// IVisStimInput is stimulus from an image
type IVisStimInput interface {
	Configure(imageFile string) error

	SetSize(string)

	EnableExpand(bool)
	SetExpand(string)
	Expand(value int) string

	SetStimulusAt(x, y int)
	GetStimulus() []int
	GetStimulusComp(value int) []int
}
