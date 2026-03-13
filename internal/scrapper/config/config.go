package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return &Config{}, fmt.Errorf("failed to load .env: %w", err)
	}
	
	port := os.Getenv("SCRAPPER_PORT")
	if port == "" {
		return nil, fmt.Errorf("SCRAPPER_PORT is not set")
	}
	
	return &Config{
		Port: port,
	}, nil
}

//доделываю