package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

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
func AppConfig(configPath string) (*Config, error) {

	if err := ValidateConfigPath(configPath); err != nil {
		log.Fatal("Config file path error:", err)
	}

	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		log.Fatal("Config file path error:", err)
	}
	if s.IsDir() {
		log.Fatal("'%s' is a directory, not a normal file", path)
	}
	return nil
}
