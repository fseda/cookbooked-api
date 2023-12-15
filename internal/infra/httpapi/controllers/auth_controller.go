package controllers

import (
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/config"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	RegisterUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Validate(c *fiber.Ctx) error
	GithubLogin(c *fiber.Ctx) error
}

type authController struct {
	service services.AuthService
	env     *config.Config
}

func NewAuthController(service services.AuthService, env *config.Config) AuthController {
	return &authController{service, env}
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

type GithubCodeRequest struct {
	Code string `json:"code"`
}

type GithubAccessTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type GithubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`      // "repo,gist"
	TokenType   string `json:"token_type"` // Bearer
}

type GithubUserResponse struct {
}

func (a *authController) GithubLogin(c *fiber.Ctx) error {

	var req GithubCodeRequest

	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError("Unable to parse body.")
	}

	token, err := a.service.GithubLogin(req.Code)
	if err != nil {
		return httpstatus.InternalServerError(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
