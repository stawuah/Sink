package sinks

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"sink/pkg/logger"
)

type ConsoleSink struct {
	pretty bool
}

func NewConsole(pretty bool) *ConsoleSink {
	return &ConsoleSink{pretty: pretty}
}

func (c *ConsoleSink) Write(e logger.Event) error {
	if c.pretty {
		return c.writePretty(e)
	}
	return c.writeJSON(e)
}

func (c *ConsoleSink) writePretty(e logger.Event) error {
	fmt.Printf("[%s] %s | %s | %s",
		e.Time.Format("15:04:05"),
		e.Level,
		e.Service,
		e.Message,
	)

	if len(e.Fields) > 0 {
		fmt.Print(" |")
		for k, v := range e.Fields {
			fmt.Printf(" %s=%v", k, v)
		}
	}
	fmt.Println()
	return nil
}

func (c *ConsoleSink) writeJSON(e logger.Event) error {
	data := map[string]any{
		"time":    e.Time.Format(time.RFC3339),
		"level":   e.Level.String(),
		"service": e.Service,
		"env":     e.Env,
		"message": e.Message,
	}

	if len(e.Fields) > 0 {
		data["fields"] = e.Fields
	}

	enc := json.NewEncoder(os.Stdout)
	return enc.Encode(data)
}

func (c *ConsoleSink) Close() error {
	return nil
}
