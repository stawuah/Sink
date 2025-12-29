package logger

import (
	"fmt"
	"sync"
	"time"
)

// Logger is the main logging interface
type Logger interface {
	Debug(msg string, fields Fields)
	Info(msg string, fields Fields)
	Warn(msg string, fields Fields)
	Error(msg string, fields Fields)
	Shutdown() error
}

// Config holds logger configuration
type Config struct {
	Service  string
	Env      string
	MinLevel Level
	Async    bool
}

type logger struct {
	cfg   Config
	sinks []Sink
	mu    sync.Mutex
}

// New creates a new logger instance
func New(cfg Config, sinks ...Sink) Logger {
	if cfg.Service == "" {
		cfg.Service = "unknown"
	}
	if cfg.Env == "" {
		cfg.Env = "development"
	}
	return &logger{
		cfg:   cfg,
		sinks: sinks,
	}
}

func (l *logger) log(level Level, msg string, fields Fields) {
	if level < l.cfg.MinLevel {
		return
	}

	event := Event{
		Time:    time.Now(),
		Level:   level,
		Service: l.cfg.Service,
		Env:     l.cfg.Env,
		Message: msg,
		Fields:  fields,
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	for _, sink := range l.sinks {
		if err := sink.Write(event); err != nil {
			fmt.Printf("sink error: %v\n", err)
		}
	}
}

func (l *logger) Debug(msg string, fields Fields) {
	l.log(DebugLevel, msg, fields)
}
func (l *logger) Info(msg string, fields Fields)  { l.log(InfoLevel, msg, fields) }
func (l *logger) Warn(msg string, fields Fields)  { l.log(WarnLevel, msg, fields) }
func (l *logger) Error(msg string, fields Fields) { l.log(ErrorLevel, msg, fields) }

func (l *logger) Shutdown() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	var errs []error
	for _, sink := range l.sinks {
		if err := sink.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("shutdown errors: %v", errs)
	}
	return nil
}
