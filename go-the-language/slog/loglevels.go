package main

import (
	"log/slog"
	"os"
)

func main() {
	// The higher the level, the more severe the event
	slog.Debug("Deep into details")
	slog.Info("Just a message")
	slog.Warn("It`s getting serious")
	slog.Error("Now thats bad")

	// Own handler, set level
	var programLevel = new(slog.LevelVar) // Info by default
	textLogger := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
    AddSource: true,

    Level: programLevel,
  })
	slog.SetDefault(slog.New(textLogger))

	programLevel.Set(slog.LevelDebug)
	slog.Debug("Can you read me now?")
	programLevel.Set(slog.LevelInfo)
	slog.Debug("Heeelloo?")
  

}
