package api

// ILogger is a general logging interface
type ILogger interface {
	LogError(string)
	LogInfo(string)
	Close()
}
