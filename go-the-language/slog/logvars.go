package main

import (
	"log/slog"
	"os"
)

func main(){

	// Logging vars
	// Define integer
	var programLevel = new(slog.LevelVar) // Info by default
	jsonLogger := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(jsonLogger))
	i := 42
	slog.Info("That does not work: ", i)
	slog.Info("Either: ", "i: ", i)
	slog.Info("Or:", slog.Int("i", i))
}