package main

import (
	"fmt"
	"os"

	"sink/pkg/logger"
	"sink/pkg/sinks"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "test" {
		runTest()
		return
	}

	fmt.Println("sink - lightweight observability logger")
	fmt.Println("usage: sink test")
}

func runTest() {
	log := logger.New(
		logger.Config{
			Service:  "test-service",
			Env:      "development",
			MinLevel: logger.DebugLevel,
		},
		sinks.NewConsole(true),
	)
	defer log.Shutdown()

	log.Info("sink initialized", nil)
	log.Debug("starting test sequence", logger.Fields{
		"version": "0.1.0",
		"mode":    "test",
	})
	log.Warn("this is a warning", logger.Fields{
		"retry_count": 3,
	})
	log.Error("simulated error", logger.Fields{
		"error_code": "E001",
		"user_id":    "u123",
	})
	log.Info("test complete", nil)
}
