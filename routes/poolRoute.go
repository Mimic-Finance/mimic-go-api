package routes

import (
	"mimic-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func PoolRoute(app *fiber.App) {
	app.Post("/pool", controllers.CreatePool)
	app.Get("/pool/:address", controllers.GetAPool)
	app.Get("/pools", controllers.GetAllPools)
}
