package config

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
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

// Reads config.yml and "unamrshals" the data.
func ReadConfig() (*Config, error) {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, fmt.Errorf("ERROR reading config file: %v", err)
	}

	cfg := Config{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("ERROR parsing config file: %v", err)
	}

	return &cfg, nil
}
