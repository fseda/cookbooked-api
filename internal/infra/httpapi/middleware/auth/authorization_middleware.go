package middlewares

import (
	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	jwtutil "github.com/fseda/cookbooked-api/internal/infra/jwt"
	"github.com/gofiber/fiber/v2"
)

func RoleRequired(requiredRole models.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwtutil.CustomClaims)
		if user.Role != string(requiredRole) {
			return httpstatus.ForbiddenError(fiber.ErrForbidden.Error())
		}

		return c.Next()
	}
}


