package main

import (
	"node/config"
	"node/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.SetupHealthRoute(app)
	router.Receive(app)
	port := config.Config("PORT")
	app.Listen("0.0.0.0:" + port)
}
