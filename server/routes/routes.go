package routes

import (
	"github.com/brianxor/nudata-solver/server/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	nudata := app.Group("/nudata")

	nudata.Post("/solve", handlers.HandleNudataSolver)

}
