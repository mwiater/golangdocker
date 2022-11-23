package main

import (
	"log"

	"github.com/mattwiater/golangdocker/api"
)

func main() {
	app := api.SetupRoute()

	log.Fatal(app.Listen(":5000"))
}
