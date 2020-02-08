package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/wdevore/Deuron8-Go/interfaces"
)

// API is the main logger for the simulation
var API interfaces.ILogger

type logg struct {
	infoFile  *os.File
	errorFile *os.File
}

// New constructs an ILog object
func New(config interfaces.IConfig) interfaces.ILogger {
	o := new(logg)

	rootPath := config.LogRoot()

	var infoFile *os.File
	var err error

	if config.ExitState() == "Paused" {
		infoFile, err = os.OpenFile(rootPath+config.InfoLogFileName(), os.O_APPEND|os.O_WRONLY, 0600)
	} else {
		infoFile, err = os.OpenFile(rootPath+config.InfoLogFileName(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
		// f, err := os.Create(rootPath + "/" + config.InfoLogFileName())
	}

	if err != nil {
		log.Fatalln(err)
	}

	var errorFile *os.File
	errorFile, err = os.OpenFile(rootPath+config.ErrLogFileName(), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalln(err)
	}

	o.infoFile = infoFile
	o.errorFile = errorFile

	o.LogInfo("All Log files have been opened")
	return o
}

// Close flushes and closes all log files. Use it with defer
func (l *logg) Close() {
	err := l.infoFile.Close()
	if err != nil {
		log.Fatalln(err)
	}
	err = l.errorFile.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func logToFile(file *os.File, msg string) {
	t := time.Now()
	now := fmt.Sprintf("%02d-%02d %02d:%02d:%02d ",
		t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	if _, err := file.WriteString(now + msg + "\n"); err != nil {
		log.Fatalln(err)
	}
}

// LogInfo logs to info log file. See config.json
func (l *logg) LogInfo(msg string) {
	logToFile(l.infoFile, msg)
}

// LogError logs to error log file. See config.json
func (l *logg) LogError(msg string) {
	logToFile(l.errorFile, "--------------------------------")
	logToFile(l.errorFile, msg)

	// We skip 1 because we don't want 'this' LogError runtime info
	// we want what "called" LogError
	// pc, file, line, _ := runtime.Caller(1)
	// rt := fmt.Sprintf("AT: %s : %s() : %d", filepath.Base(file), filepath.Base(runtime.FuncForPC(pc).Name()), line)
	pc, _, line, _ := runtime.Caller(1)
	rt := fmt.Sprintf("AT: %s() : %d", filepath.Base(runtime.FuncForPC(pc).Name()), line)

	logToFile(l.errorFile, rt)
}
