package main

import (
	"log/slog"
	"os"
)

func main() {
	
	// Own handler, set level
	var programLevel = new(slog.LevelVar) // Info by default
	//create new io writer with file goa.log
	file, err := os.OpenFile("goa.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("Cant open file, here is why: ", "Error", err)
		os.Exit(1)
	}
	defer file.Close()
	textLogger := slog.NewTextHandler(file, &slog.HandlerOptions{
    	AddSource: true,
    	Level: programLevel,
  	})
	slog.SetDefault(slog.New(textLogger))
	slog.Info("Hello file?")
  

}
