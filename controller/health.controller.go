package controller

import "github.com/gofiber/fiber/v2"

func Health(c *fiber.Ctx) error{
	return c.Status(200).JSON("OK")
}
