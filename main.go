package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mwiater/golangdocker/api"
	"github.com/mwiater/golangdocker/config"
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":" + cfg["SERVERPORT"]); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
	// Complete any cleanup tasks here...
}
