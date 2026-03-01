package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	TelegramToken string
	LogLevel string
}


func Load() (Config, error) {
	_ = godotenv.Load()
	token := os.Getenv("APP_TELEGRAM_TOKEN");
	log_lvl := os.Getenv("LOG_LEVEL")
	if log_lvl == "" {
		log_lvl = "info"
	}

	if token == "" {
		return Config{}, fmt.Errorf("APP_TELEGRAM_TOKEN is not get");
	}

	return Config {
		TelegramToken: token,
		LogLevel: log_lvl,
	}, nil
}


func ParseLogLevel(log_lvl string) slog.Level {
	switch (log_lvl) {
	case "info":
		return slog.LevelInfo
	case "debug":
		return slog.LevelDebug
	case "error":
		return slog.LevelError
	default:
        return slog.LevelInfo
    }
}
