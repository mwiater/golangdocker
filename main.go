package main

import (
	"embed"
	"log"
	"strconv"
	//"fmt"

	"github.com/mattwiater/golangdocker/api"
	"github.com/mattwiater/golangdocker/config"
)

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
