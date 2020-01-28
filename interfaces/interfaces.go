package interfaces

// ILogger is a general logging interface
type ILogger interface {
	LogError(string)
	LogInfo(string)
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
