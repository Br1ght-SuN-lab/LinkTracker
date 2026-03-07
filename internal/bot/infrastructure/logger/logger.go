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


func GetLevel(s string) LogLevel {
	switch LogLevel(s) {
	case LogInfo, LogDebug, LogWarn, LogError:
		return LogLevel(s)
	default:
		return LogInfo
	}
}


func (l LogLevel) ToSlog() slog.Level {
    switch l {
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


func New(logLevel string) (logger *slog.Logger) {
	level := GetLevel(logLevel);
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level.ToSlog(),
	}))
	return logger
}