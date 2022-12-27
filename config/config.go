package config

import (
	"bufio"
	"embed"
	"strings"
)

// AppConfig returns a new decoded Config struct
func AppConfig(fs embed.FS) (map[string]string, error) {
	var envs map[string]string
	envs = make(map[string]string)
	envsFromFile, _ := fs.ReadFile(".env")

	scanner := bufio.NewScanner(strings.NewReader(string(envsFromFile)))
	for scanner.Scan() {
		env := strings.Split(scanner.Text(), "=")
		envs[env[0]] = env[1]
	}

	return envs, nil
}
