package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mattwiater/golangdocker/api"
)

func main() {

	app := api.SetupRoute()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go Fiber API")
	})

	// err := app.ListenTLS(":5000", "./certs/192.168.0.99.crt", "./certs/192.168.0.99.key")
	err := app.Listen(":5000")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
