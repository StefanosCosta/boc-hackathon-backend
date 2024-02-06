package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var (
	ClientID     string
	ClientSecret string
	BaseURI      string
)

type Config struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	BaseURI      string `yaml:"base_uri"`
}

func LoadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return err
	}

	ClientID = config.ClientID
	ClientSecret = config.ClientSecret
	BaseURI = config.BaseURI

	return nil
}
