package config

import (
	"fmt"
	"os"
	"log/slog"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type LogLevel string 

const (
	LogDebug LogLevel = "debug"
	LogInfo  LogLevel = "info"
	LogWarn  LogLevel = "warn"
	LogError LogLevel = "error"
)

func ParseLogLevel(s string) LogLevel {
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

type Config struct {
	TelegramToken string
	LogLevel string `yaml:"loglevel"`
}


func Load(path string) (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("failed to load .env: %w", err)
	}

	cfg := Config{}

	b, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read %s: %w", path, err)
	}
	
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse %s: %w", path, err)
	}

	cfg.TelegramToken = os.Getenv("APP_TELEGRAM_TOKEN");
	if cfg.TelegramToken == "" {
		return Config{}, fmt.Errorf("APP_TELEGRAM_TOKEN is not get");
	}

	if cfg.LogLevel == "" {
		cfg.LogLevel = "info" //default value
	}

	return cfg, nil
}
