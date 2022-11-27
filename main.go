package main

import (
	"embed"
	"log"
	"strconv"

	"github.com/mattwiater/golangdocker/api"
	"github.com/mattwiater/golangdocker/config"
)

//go:embed config/appConfig.yml
var f embed.FS

func main() {
	cfg, err := config.AppConfig("./config/appConfig.yml")
	if err != nil {
		log.Fatal(err)
	}

	app := api.SetupRoute(cfg)

	log.Fatal(app.Listen(":" + strconv.Itoa(cfg.Server.Port)))
}
