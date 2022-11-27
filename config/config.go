package config

import (
	"gopkg.in/yaml.v3"
)

var cfg Config

// Config struct for webapp config
type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Options struct {
		Debug bool `yaml:"debug"`
	} `yaml:"options"`
}

// AppConfig returns a new decoded Config struct
func AppConfig(configData []byte) (Config, error) {

	err := yaml.Unmarshal(configData, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
