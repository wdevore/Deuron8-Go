package log

import (
	"fmt"
	"github.com/wdevore/Deuron8-Go/interfaces"
)

// Logger implements ILog
type Logger struct {
	infoLogFile  string
	errorLogFile string
}

// New constructs an ILog object
func New(config interfaces.IConfig) interfaces.ILogger {
	o := new(Logger)
	o.infoLogFile = config.InfoLogFileName()
	return o
}

// LogInfo logs to info log file. See config.json
// Implements ILog interface
func (l *Logger) LogInfo(msg string) {
	fmt.Println(msg)
}

// LogError logs to error log file. See config.json
// Implements ILog interface
func (l *Logger) LogError(msg string) {
	fmt.Println(msg)
}
