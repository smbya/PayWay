package config

import (
	"log/slog"
	"os"
)

func createLogger() *slog.Logger {
	level := getLogLevelFromEnv()

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	return slog.New(handler)
}

func getLogLevelFromEnv() slog.Level {
	switch getEnvWithDefault("LOGGER_LEVEL", "info") {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
