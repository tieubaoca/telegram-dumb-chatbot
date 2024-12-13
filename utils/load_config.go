package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port        int    `yaml:"port"`
	LogFile     string `yaml:"log_file"`
	ErrorFile   string `yaml:"error_file"`
	WebhookURI  string `yaml:"webhook_uri"`
	LlmEndpoint string `yaml:"llm_endpoint"`
	SDEndpoint  string `yaml:"sd_endpoint"`
}

func LoadConfig(configFile string) (*Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
