package config 

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)


type Config struct {
	TelegramToken string
}


func Load() (Config, error) {
	_ = godotenv.Load()
	token := os.Getenv("APP_TELEGRAM_TOKEN");

	if token == "" {
		return Config{}, fmt.Errorf("APP_TELEGRAM_TOKEN is not get");
	}

	return Config {
		TelegramToken: token,
	}, nil
}