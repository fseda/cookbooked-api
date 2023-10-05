package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fseda/cookbooked-api/internal/domain/models"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/validation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	*models.User
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
		return httpstatus.BadRequestError(strings.Join(errMsgs, " and "))
	}

	user, err := u.service.Create(req.Username, req.Email, req.Password)
	if err != nil {
		if err == u.service.CreateUserErrors.EmailExists {
			msg := fmt.Sprintf("%s (%s)", u.service.CreateUserErrors.EmailExists.Error(), req.Email)
			return httpstatus.BadRequestError(msg)
		}

		if err == u.service.CreateUserErrors.UsernameExists {
			msg := fmt.Sprintf("%s (%s)", u.service.CreateUserErrors.UsernameExists.Error(), req.Username)
			return httpstatus.BadRequestError(msg)
		}

		return httpstatus.InternalServerError(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(CreateUserResponse{
		user,
	})
}

func (u *UserController) FindOne(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return httpstatus.BadRequestError("Invalid id. Should be a positive integer.")
	}

	user, err := u.service.FindByID(uint(id))
	if err == gorm.ErrRecordNotFound {
		msg := fmt.Sprintf("User with ID %d not found", id)
		return httpstatus.NotFoundError(msg)
	}
	if err != nil {
		return httpstatus.InternalServerError(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (u *UserController) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return httpstatus.BadRequestError("Invalid id. Should be a positive integer.")
	}

	ra, err := u.service.Delete(uint(id))
	if err != nil {
		return httpstatus.InternalServerError("Could not delete user.")
	}

	if ra == 0 {
		return c.SendStatus(fiber.StatusNoContent)
	}

	return c.SendStatus(fiber.StatusOK)
}


