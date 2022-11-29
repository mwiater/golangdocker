package main

import (
	"embed"
	"log"
	"strconv"

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

//go:embed config/appConfig.yml
var conf embed.FS

func main() {
	configData, _ := conf.ReadFile("config/appConfig.yml")

	cfg, err := config.AppConfig(configData)
	if err != nil {
		log.Fatal(err)
	}

	app := api.SetupRoute(cfg)

	log.Fatal(app.Listen(":" + strconv.Itoa(cfg.Server.Port)))
}
