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

	err := app.Listen(":5000")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
