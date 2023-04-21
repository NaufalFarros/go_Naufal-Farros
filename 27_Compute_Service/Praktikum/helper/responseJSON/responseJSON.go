package responseJSON

import (
	"github.com/gofiber/fiber/v2"
)

func ResponseJSON(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}
