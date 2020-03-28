package internal

import (
	"log"
	"os"
)

var (
	logFlags    = log.Ldate | log.Ltime
	info        = log.New(os.Stdout, "INFO: ", logFlags)
	errorLog    = log.New(os.Stderr, "ERROR: ", logFlags)
	warning     = log.New(os.Stderr, "WARNING: ", logFlags)
)

// DebugLogger provides an interface to call the logging functions
type DebugLogger interface {
	LogInfo(v ...interface{})
	LogError(v ...interface{})
	LogWarning(v ...interface{})
	PrintVersion(v ...interface{})
}

// Debug contains information about the debug level
type Debug struct {
	DebugLevel  bool
}

func (d Debug) LogInfo(v ...interface{}) {
	if d.DebugLevel {
		info.Println(v...)
	}
}

func (d Debug) LogError(v ...interface{}) {
	errorLog.Println(v...)
}

func (d Debug) LogWarning(v ...interface{}) {
	if d.DebugLevel {
		warning.Println(v...)
	}
}

func (d Debug) PrintVersion(v ...interface{}) {
	info.Println(v...)
}
