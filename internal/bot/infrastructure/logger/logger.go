package logger

import (
	"log/slog"
	"os"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/config"
)


func New(log_lvl string) (logger *slog.Logger) {
	level := config.ParseLogLevel(log_lvl).ToSlog()
	logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	return logger
}