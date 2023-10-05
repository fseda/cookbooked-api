package controllers

import (
	"fmt"
	"strconv"
	"strings"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/models"
	modelvalidation "github.com/fseda/cookbooked-api/internal/domain/models/validation"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/validation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController interface {
	RegisterUser(c *fiber.Ctx) error
	FindOne(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{service}
}

type createUserRequest struct {
	Username string `json:"username" validate:"required=true,min=3,max=255"`
	Email    string `json:"email" validate:"required=true,email"`
	Password string `json:"password" validate:"required=true,min=6,max=72"`
}

type createUserResponse struct {
	*models.User
}

func (u *userController) RegisterUser(c *fiber.Ctx) error {
	var req createUserRequest
	if err := c.BodyParser(&req); err != nil {
		return httpstatus.UnprocessableEntityError("Unable to parse request body.")
	}

	if errMsgs := validation.MyValidator.CreateErrorResponse(req); len(errMsgs) > 0 {
		return httpstatus.BadRequestError(strings.Join(errMsgs, " and "))
	}

	if modelvalidation.IsEmailLike(req.Username) {
		return httpstatus.BadRequestError(globalerrors.UserInvalidUsername.Error())
	}

	user, err := u.service.Create(req.Username, req.Email, req.Password)
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

	return c.Status(fiber.StatusCreated).JSON(createUserResponse{
		user,
	})
}

func (u *userController) FindOne(c *fiber.Ctx) error {
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

func (u *userController) Delete(c *fiber.Ctx) error {
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
