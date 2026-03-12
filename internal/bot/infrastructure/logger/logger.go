package logger

import (
	"log/slog"
	"os"
)

type LogLevel string

const (
	LogDebug LogLevel = "debug"
	LogInfo  LogLevel = "info"
	LogWarn  LogLevel = "warn"
	LogError LogLevel = "error"
)

func GetLevel(s LogLevel) slog.Level {
	switch LogLevel(s) {
	case LogDebug:
		return slog.LevelDebug
	case LogInfo:
		return slog.LevelInfo
	case LogWarn:
		return slog.LevelWarn
	case LogError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}


func New(logLevel LogLevel) (logger *slog.Logger) {
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: GetLevel(logLevel),
	}))
	return logger
}
