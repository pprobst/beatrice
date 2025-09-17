package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Struct to hold the configuration info defined in config.yml.
type Config struct {
	Title       string
	Description string
	Author      string
	Footer      string
	Theme       string
	Posts       string
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Title == "" {
		return fmt.Errorf("title is required in config.yml")
	}
	if c.Author == "" {
		return fmt.Errorf("author is required in config.yml")
	}
	if c.Theme != "light" && c.Theme != "dark" {
		return fmt.Errorf("theme must be 'light' or 'dark', got '%s'", c.Theme)
	}
	return nil
}

// ReadConfig reads the configuration file and unmarshals the data.
func ReadConfig() (*Config, error) {
	return ReadConfigFromFile("config.yml")
}

// ReadConfigFromFile reads a configuration file from the specified path.
func ReadConfigFromFile(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %v", configPath, err)
	}

	cfg := Config{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %v", err)
	}

	return &cfg, nil
}
