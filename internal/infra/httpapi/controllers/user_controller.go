package controllers

import (
	"strconv"
	"strings"

	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/validation"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service}
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required=true,min=3,max=255,alphanum"`
	Email    string `json:"email" validate:"required=true,email"`
	Password string `json:"password" validate:"required=true,min=6,max=72"`
}

type CreateUserResponse struct {
	ID uint `json:"id"`
}

func (u *UserController) Create(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	if errMsgs := validation.MyValidator.CreateErrorResponse(req); len(errMsgs) > 0 {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		}
	}

	id, err := u.service.Create(req.Username, req.Email, req.Password)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	return c.Status(fiber.StatusCreated).JSON(CreateUserResponse{
		ID: id,
	})
}

func (u *UserController) FindOne(c *fiber.Ctx) error {
	idStr := c.Params("id", "0")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return httpstatus.BadRequestError("Invalid id. Should be an integer.")
	}

	user, err := u.service.FindByID(uint(id))
	if err != nil {
		return httpstatus.NotFoundError(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
