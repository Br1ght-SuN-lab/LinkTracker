package config 

import (
	"fmt"
	"os"
)


type Config struct {
	TelegramToken string
}


func Load() (Config, error) {
	token := os.Getenv("APP_TELEGRAM_TOKEN");

	if token == "" {
		return Config{}, fmt.Errorf("APP_TELEGRAM_TOKEN is not get");
	}

	return Config {
		TelegramToken: token,
	}, nil
}