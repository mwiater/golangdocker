// Package config implements .env file application configuration
package config

import (
	"embed"
	"strings"

	"github.com/mattwiater/golangdocker/common"
)

var EnvVarsFile embed.FS

// AppConfig returns a new decoded Config struct
func AppConfig() (map[string]string, error) {
	envVars, _ := EnvVarsFile.ReadFile(".env")

	lines := common.SplitStringLines(string(envVars))

	var envs = make(map[string]string)
	for _, line := range lines {
		keyValuePair := strings.Split(line, "=")
		envs[keyValuePair[0]] = keyValuePair[1]
	}
	return envs, nil
}
