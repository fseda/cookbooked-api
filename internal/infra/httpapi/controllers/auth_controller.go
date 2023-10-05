package controllers

import (
	"strings"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/validation"
	jwtutil "github.com/fseda/cookbooked-api/internal/infra/jwt"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	Profile(c *fiber.Ctx) error
}

type authController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) AuthController {
	return &authController{service}
}

type loginUserRequest struct {
	Username string `json:"username" validate:"required=true"`
	Password string `json:"password" validate:"required=true"`
}

type loginUserResponse struct {
	Token string `json:"token"`
}

type userProfileResponse struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (a *authController) Login(c *fiber.Ctx) error {
	var req loginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError("Unable to parse body.")
	}

	if errMsgs := validation.MyValidator.CreateErrorResponse(req); len(errMsgs) > 0 {
		return httpstatus.BadRequestError(strings.Join(errMsgs, " and "))
	}

	token, err := a.service.Login(req.Username, req.Password)
	if err != nil {
		switch err {
		case globalerrors.GlobalInternalServerError:
			return httpstatus.InternalServerError(err.Error())
		case globalerrors.AuthInvalidCredentials:
			return httpstatus.UnauthorizedError(err.Error())
		default:
			return httpstatus.InternalServerError("Something went wrong. Please try again later.")
		}
	}

	return c.Status(fiber.StatusOK).JSON(loginUserResponse{
		token,
	})
}

func (a *authController) Profile(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*jwtutil.CustomClaims)

	userID := claims.UserID
	username := claims.Username
	role := claims.Role

	return c.Status(fiber.StatusOK).JSON(userProfileResponse{
		UserID: userID,
		Username: username,
		Role: role,
	})
}
