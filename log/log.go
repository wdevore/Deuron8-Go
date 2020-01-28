package log

import (
	"fmt"
	"github.com/wdevore/Deuron8-Go/interfaces"
)

// Logger implements ILog
type Logger struct {
	infoLogPath string
}

// New constructs an ILog object
func New() interfaces.ILogger {
	o := new(Logger)
	o.infoLogPath = "info2.log"
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
