package middlewares

import (
	"strconv"

	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/gofiber/fiber/v2"
)

func ValidateID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if _, err := strconv.ParseUint(c.Params("id"), 10, 64); err != nil {
			return httpstatus.BadRequestError("Invalid ID")
		}

		return c.Next()
	}
}
