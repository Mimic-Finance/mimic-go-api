package main

import (
	"mimic-api/configs"
	"mimic-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	//run database
	configs.ConnectDB()

	//routes
	routes.PoolRoute(app)

	app.Listen(":3000")
}
