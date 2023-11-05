package router

import (
	"node/controller"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func Receive(app *fiber.App) {
	client := app.Group("/")
	client.Post("/receive", controller.Receive)
}
