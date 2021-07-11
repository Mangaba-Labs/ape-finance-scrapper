package handler

import "github.com/gofiber/fiber/v2"

// HealthCheck handler
func HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"status": "healthy"})
}
