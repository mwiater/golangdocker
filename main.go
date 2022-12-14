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
var fs embed.FS

func main() {

	cfg, err := config.AppConfig(fs)
	if err != nil {
		log.Fatal("Error: config.AppConfig()")
	}

	app := api.SetupRoute(cfg)

	log.Fatal(app.Listen(":" + cfg["SERVERPORT"]))
}
