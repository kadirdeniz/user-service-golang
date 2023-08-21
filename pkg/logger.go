package pkg

import (
	"log/slog"
	"os"
)

func NewLogger() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}
