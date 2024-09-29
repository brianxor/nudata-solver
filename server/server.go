package server

import (
	"fmt"
	"github.com/brianxor/nudata-solver/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start(serverHost string, serverPort string) error {
	appConfig := fiber.Config{
		AppName: "NuData Solver",
	}

	app := fiber.New(appConfig)

	app.Use(logger.New())
	
	routes.SetupRoutes(app)

	serverAddr := fmt.Sprintf("%s:%s", serverHost, serverPort)

	if err := app.Listen(serverAddr); err != nil {
		return err
	}

	return nil
}
