package generator

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	API        string   `yaml:"api"`
	Version    string   `yaml:"version"`
	Extensions []string `yaml:"extensions"`
	Platforms  []string `yaml:"platforms,omitempty"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}