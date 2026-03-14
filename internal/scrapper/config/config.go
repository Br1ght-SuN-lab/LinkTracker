package config

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port string
}

func Load(path string) (*Config, error) {
	cfg := &Config{}
	file, err := os.ReadFile(path)
	
	if err != nil {
		return &Config{}, fmt.Errorf("read %s: %w", path, err)
	}

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return &Config{}, fmt.Errorf("parse %s: %w", path, err)
	}

	return cfg, nil
}