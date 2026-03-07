package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	TelegramToken string
	LogLevel string `yaml:"loglevel"`
}


func Load(path string) (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return &Config{}, fmt.Errorf("failed to load .env: %w", err)
	}

	cfg := Config{}

	file, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, fmt.Errorf("read %s: %w", path, err)
	}
	
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return &Config{}, fmt.Errorf("parse %s: %w", path, err)
	}

	cfg.TelegramToken = os.Getenv("APP_TELEGRAM_TOKEN");
	if cfg.TelegramToken == "" {
		return &Config{}, fmt.Errorf("APP_TELEGRAM_TOKEN is not get");
	}

	if cfg.LogLevel == "" {
		cfg.LogLevel = "info" //default value
	}

	return &cfg, nil
}
