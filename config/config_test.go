package config_test

import (
	"embed"
	"fmt"

	"github.com/mattwiater/golangdocker/config"
)

//go:generate bash copy_env_file.sh

//go:embed .env
var fs embed.FS

func ExampleAppConfig() {
	fmt.Println(config.AppConfig(fs))
	// Output:
	// map[DEBUG:false DOCKERIMAGE:mattwiater/golangdocker DOCKERPORT:5000 SERVERPORT:5000] <nil>
}
