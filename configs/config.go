package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Redis struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"redis"`
	Interval time.Duration `yaml:"interval"`
	Limit    int64         `yaml:"limit"`
	UserID   int64         `yaml:"userID"`
}

func InitConfig(filename string) (*Config, error) {
	var config Config

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while reading the configurations: %w", err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while parsing the yaml file: %w", err)
	}
	if config.Limit <= 0 {
		config.Limit = 1
	}
	if config.Interval <= 0 {
		config.Interval = 10
	}
	if config.Redis.Host == "" {
		config.Redis.Host = "local"
	}
	if config.Redis.Port == "" {
		config.Redis.Host = "6379"
	}

	return &config, nil
}
