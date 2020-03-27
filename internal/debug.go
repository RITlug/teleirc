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
	LogInfo(string)
	LogError(error)
	LogWarning(string)
	PrintVersion(string)
}

// Debug contains information about the debug level
type Debug struct {
	DebugLevel  bool
}

func (d Debug) LogInfo(message string) {
	if d.DebugLevel {
		info.Println(message)
	}
}

func (d Debug) LogError(message error) {
	errorLog.Println(message)
}

func (d Debug) LogWarning(message string) {
	if d.DebugLevel {
		warning.Println(message)
	}
}

func (d Debug) PrintVersion(message string) {
	info.Println(message)
}
