package logger

import (
	"log/slog"
	"os"
)

// New creates a structured logger for the API service.
func New(level string) *slog.Logger {
	var slogLevel slog.Level
	if level == "debug" {
		slogLevel = slog.LevelDebug
	} else {
		slogLevel = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slogLevel})
	return slog.New(handler)
}
