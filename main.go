package main

import (
	"embed"
	"log"

	"github.com/mattwiater/golangdocker/api"
	"github.com/mattwiater/golangdocker/config"
)

// @title Golang Docker
// @version 1.0
// @description Boilerplate Fiber Rest API for use in Docker

// @contact.name Matt J. Wiater
// @contact.url https://github.com/mwiater

// @license.name MIT License
// @license.url https://en.wikipedia.org/wiki/MIT_License

// @BasePath /

//go:embed .env
var envVarsFile embed.FS

func main() {
	config.EnvVarsFile = envVarsFile

	cfg, err := config.AppConfig()
	if err != nil {
		log.Fatal("Error: config.AppConfig()")
	}

	app := api.SetupApi()
	log.Fatal(app.Listen(":" + cfg["SERVERPORT"]))
}
