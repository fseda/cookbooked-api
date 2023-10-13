package middlewares

import (
	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/gofiber/fiber/v2"
)

func ValidateID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idInt, err := c.ParamsInt("id")
		if err != nil || idInt <= 0{
			return httpstatus.BadRequestError(globalerrors.GlobalInvalidID.Error())
		}

		return c.Next()
	}
}
