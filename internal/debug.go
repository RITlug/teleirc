package internal

import (
	"log"
	"os"
)

var (
	logFlags = log.Ldate | log.Ltime
	info     = log.New(os.Stdout, "INFO: ", logFlags)
	debug    = log.New(os.Stdout, "DEBUG: ", logFlags)
	errorLog = log.New(os.Stderr, "ERROR: ", logFlags)
	warning  = log.New(os.Stderr, "WARNING: ", logFlags)
	plain    = log.New(os.Stdout, "", 0)
)

// DebugLogger provides an interface to call the logging functions
type DebugLogger interface {
	LogInfo(v ...interface{})
	LogDebug(v ...interface{})
	LogError(v ...interface{})
	LogWarning(v ...interface{})
	PrintVersion(v ...interface{})
}

// Debug contains information about the debug level
type Debug struct {
	DebugLevel bool
}

// LogInfo prints info-level messages to standard out
func (d Debug) LogInfo(v ...interface{}) {
	info.Println(v...)
}

// LogDebug prints debug-level messages to standard out
func (d Debug) LogDebug(v ...interface{}) {
	if d.DebugLevel {
		debug.Println(v...)
	}
}

// LogError prints error-level messages to standard out
func (d Debug) LogError(v ...interface{}) {
	errorLog.Println(v...)
}

// LogWarning prints warning-level messages to standard out
func (d Debug) LogWarning(v ...interface{}) {
	if d.DebugLevel {
		warning.Println(v...)
	}
}

// PrintVersion prints the TeleIRC version number
func (d Debug) PrintVersion(v ...interface{}) {
	plain.Println(v...)
}
