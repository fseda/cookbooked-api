package controllers

import (
	"fmt"
	"strings"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/validation"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	RegisterUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
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

//	@Summary		Login user into the app
//	@Description	Logs an existing user into the app
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user-credentials	body		loginUserRequest	true	"User credentials"
//	@Success		200					{object}	loginUserResponse
//	@Router			/auth/login [post]
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

type registerUserRequest struct {
	Username string `json:"username" validate:"required=true,min=3,max=255"`
	Email    string `json:"email" validate:"required=true,email"`
	Password string `json:"password" validate:"required=true,min=6,max=72"`
}

type registerUserResponse struct {
	Token string `json:"token"`
}

//	@Summary		Register user in the app
//	@Description	Registers a new user in the app
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user-info	body		registerUserRequest		true	"New user credentials"
//	@Success		201			{object}	registerUserResponse	"User created successfully, returns jwt token"
//	@Router			/auth/signup [post]
func (a *authController) RegisterUser(c *fiber.Ctx) error {
	var req registerUserRequest
	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError("Unable to parse request body.")
	}

	if errMsgs := validation.MyValidator.CreateErrorResponse(req); len(errMsgs) > 0 {
		return httpstatus.BadRequestError(strings.Join(errMsgs, " and "))
	}

	user, err := a.service.Create(req.Username, req.Email, req.Password)
	if err != nil {
		if err == globalerrors.UserEmailExists {
			msg := fmt.Sprintf("%s (%s)", globalerrors.UserEmailExists, req.Email)
			return httpstatus.BadRequestError(msg)
		}

		if err == globalerrors.UserUsernameExists {
			msg := fmt.Sprintf("%s (%s)", globalerrors.UserUsernameExists, req.Username)
			return httpstatus.BadRequestError(msg)
		}

		return httpstatus.InternalServerError(err.Error())
	}

	token, err := a.service.Login(user.Username, req.Password)
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

	return c.Status(fiber.StatusCreated).JSON(loginUserResponse{
		Token: token,
	})
}
