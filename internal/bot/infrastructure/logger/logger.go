package logger

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/types"
	"log/slog"
	"os"
)

func GetLevel(s types.LogLevel) types.LogLevel {
	switch types.LogLevel(s) {
	case types.LogInfo, types.LogDebug, types.LogWarn, types.LogError:
		return types.LogLevel(s)
	default:
		return types.LogInfo
	}
}

func ToSlog(l types.LogLevel) slog.Level {
	switch l {
	case types.LogDebug:
		return slog.LevelDebug
	case types.LogInfo:
		return slog.LevelInfo
	case types.LogWarn:
		return slog.LevelWarn
	case types.LogError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func New(logLevel types.LogLevel) (logger *slog.Logger) {
	level := GetLevel(logLevel)
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: ToSlog(level),
	}))
	return logger
}
