package main

import (
	"log/slog"
	"os"
)

func main() {
	state := "Running"

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	statusLogger := logger.With("State:", state)

	statusLogger.Info("Hi")
	state = "Sleeping"

	statusLogger.Info("Hello again")
  statusLogger = logger.With("State:", state)

  statusLogger.Info("I am back")

}
