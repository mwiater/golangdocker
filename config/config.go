// Package config implements .env file application configuration
package config

import (
	"bufio"
	"embed"
	"strings"
)

var EnvVarsFile embed.FS

// AppConfig returns a new decoded Config struct
func AppConfig() (map[string]string, error) {
	envVars, _ := EnvVarsFile.ReadFile(".env")

	lines := splitLines(string(envVars))

	var envs = make(map[string]string)
	for _, line := range lines {
		keyValuePair := strings.Split(line, "=")
		envs[keyValuePair[0]] = keyValuePair[1]
	}
	return envs, nil
}

func splitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
