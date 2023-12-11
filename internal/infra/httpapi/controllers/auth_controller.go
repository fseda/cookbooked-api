package controllers

import (
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/gofiber/fiber/v2"
)

var singleton = false
var controller = interface{}(nil)

type AuthController interface {
	RegisterUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Validate(c *fiber.Ctx) error
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

// @Summary		Login user into the app
// @Description	Logs an existing user into the app
// @Tags			Users
// @Accept			json
// @Param			user-credentials	body	loginUserRequest	true	"User credentials"
// @Success		200
// @Header			200	{string}	Authorization	"Bearer <token>"
// @Router			/auth/login [post]
func (a *authController) Login(c *fiber.Ctx) error {
	var req loginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError("Unable to parse body.")
	}

	token, validation, err := a.service.Login(req.Username, req.Password)
	if err != nil {
		return httpstatus.InternalServerError(err.Error())
	}
	if validation.HasErrors() {
		return c.Status(fiber.StatusBadRequest).JSON(validation)
	}

	c.Set("Authorization", "Bearer "+token)

	return httpstatus.OK("OK")
}

type registerUserRequest struct {
	Username string `json:"username" validate:"required=true,min=3,max=255"`
	Email    string `json:"email" validate:"required=true,email"`
	Password string `json:"password" validate:"required=true,min=4,max=72"`
}

// @Summary		Register user in the app
// @Description	Registers a new user in the app
// @Tags			Users
// @Accept			json
// @Param			user-info	body	registerUserRequest	true	"New user credentials"
// @Success		201
// @Header			201	{string}	Authorization	"Bearer <token>"
// @Router			/auth/signup [post]
func (a *authController) RegisterUser(c *fiber.Ctx) error {
	var req registerUserRequest
	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError("Unable to parse request body.")
	}

	user, validation, err := a.service.Create(req.Username, req.Email, req.Password)
	if err != nil {
		return httpstatus.InternalServerError(err.Error())
	}
	if validation.HasErrors() {
		return c.Status(fiber.StatusBadRequest).JSON(validation)
	}

	token, _, err := a.service.Login(user.Username, req.Password)
	if err != nil {
		return httpstatus.InternalServerError(err.Error())
	}

	c.Set("Authorization", "Bearer "+token)

	return httpstatus.Created("created")
}

// @Summary		Validate JWT
// @Description	Validate the JWT provided in the Authorization header
// @Tags			auth
// @Produce		json
// @Security		ApiKeyAuth
// @Success		200
// @Failure		401
// @Router			/auth/validate [get]
func (a *authController) Validate(c *fiber.Ctx) error {
	// middleware implementation
	// if it gets to here then the token is valid

	return c.JSON(fiber.Map{
		"valid": true,
	})
}
