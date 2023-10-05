package controllers

import (
	"fmt"
	"strconv"

	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController interface {
	FindOne(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{service}
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
