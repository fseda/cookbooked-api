package middlewares

import (
	"strings"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	jwtutil "github.com/fseda/cookbooked-api/internal/infra/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTAuthMiddleware(secretKey []byte) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return httpstatus.UnauthorizedError(globalerrors.AuthMissingAuthHeader.Error())
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			return httpstatus.UnauthorizedError(globalerrors.AuthInvalidAuthHeaderFormat.Error())
		}
		token := splitToken[1]

		// Validate the token
		claims, err := jwtutil.ValidateToken(token, secretKey)
		if err != nil {
			return httpstatus.UnauthorizedError(globalerrors.AuthInvalidToken.Error())
		}

		c.Locals("user", claims)

		return c.Next()
	}
}

func ValidateJWT(secretKey []byte) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	
		_, err := jwtutil.ValidateToken(tokenStr, secretKey)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"valid": false,
			})
		}
	
		return c.Next()
	}
}
