// pkg/logger/event.go
package logger

import "time"

// Level represents log severity
type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Event is the core log structure
type Event struct {
	Time    time.Time
	Level   Level
	Service string
	Env     string
	Message string
	Fields  Fields
}

// Fields holds structured log data
type Fields map[string]any