package config

import (
	"log"

	"github.com/joho/godotenv"
)

var cfg Config

// Config struct for webapp config
type Config struct {
	Port  int
	Debug bool
}

// AppConfig returns a new decoded Config struct
func AppConfig() (map[string]string, error) {
	var appEnvs map[string]string
	appEnvs, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return appEnvs, nil
}
