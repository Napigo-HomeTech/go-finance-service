package controllers

import "github.com/gofiber/fiber/v2"

func HealthCheckHandler(c *fiber.Ctx) error {
	c.WriteString("Server online")
	return nil
}
