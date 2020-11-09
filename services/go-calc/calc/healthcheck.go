package calc

import "github.com/gofiber/fiber/v2"

// HealthCheck provides simple json with current service status
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(Response{
		Status: "OK",
		Msg:    "This service is currently running.",
		Data:   true,
	})
}
