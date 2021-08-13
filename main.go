package main

import (
	"erbeyinn/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	api := app.Group("/api")

	routes.Route(api.Group("/books"))

	app.Listen(":3000")

}
