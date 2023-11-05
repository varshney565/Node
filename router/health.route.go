package router

import (
	"node/controller"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupHealthRoute(app *fiber.App) {
	client := app.Group("/")
	client.Head("/alive", controller.Health)
}
