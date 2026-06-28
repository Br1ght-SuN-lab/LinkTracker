package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port        string `yaml:"port"`
	GithubToken string `yaml:"github_token"`
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
