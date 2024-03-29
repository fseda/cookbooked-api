package controllers

import (
	"fmt"
	"strconv"

	globalerrors "github.com/fseda/cookbooked-api/internal/domain/errors"
	"github.com/fseda/cookbooked-api/internal/domain/services"
	"github.com/fseda/cookbooked-api/internal/infra/httpapi/httpstatus"
	jwtutil "github.com/fseda/cookbooked-api/internal/infra/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController interface {
	GetOneByID(c *fiber.Ctx) error
	ExistsOne(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Profile(c *fiber.Ctx) error
	UserExists(c *fiber.Ctx) error
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{service}
}

type userProfileResponse struct {
	UserID   uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar_url"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
	Role     string `json:"role"`
}

//	 Profile retrieves the profile details of the authenticated user.
//	@Summary		Get logged in user's profile
//	@Description	Retrieve the profile of the currently authenticated user.
//	@ID				get-user-profile
//	@Tags			Users
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Success		200	{object}	userProfileResponse
//	@Failure		500	{object}	httpstatus.GlobalErrorHandlerResp	"Internal server error"
//	@Failure		401	{object}	httpstatus.GlobalErrorHandlerResp	"Unauthorized"
//	@Failure		404	{object}	httpstatus.GlobalErrorHandlerResp	"User not found"
//	@Router			/me [get]
func (u *userController) Profile(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(*jwtutil.CustomClaims)

	user, err := u.service.FindByID(userClaims.UserID)
	if err == gorm.ErrRecordNotFound {
		msg := fmt.Sprintf("User with ID %d not found", userClaims.UserID)
		return httpstatus.NotFoundError(msg)
	}
	if err != nil {
		return httpstatus.InternalServerError(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(userProfileResponse{
		UserID:   userClaims.UserID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Avatar:   user.Avatar,
		Bio:      user.Bio,
		Location: user.Location,
		Role:     userClaims.Role,
	})
}

//	 GetOneByID retrieves the details of a user by their ID.
//	@Summary		Get user by ID
//	@Description	Retrieve detailed information of a user based on their ID.
//	@ID				get-user-by-id
//	@Tags			Users
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	userProfileResponse
//	@Failure		400	{object}	httpstatus.GlobalErrorHandlerResp	"User not Found"
//	@Failure		500	{object}	httpstatus.GlobalErrorHandlerResp	"Internal Server Error"
//	@Failure		401	{object}	httpstatus.GlobalErrorHandlerResp	"Unauthorized
//	@Router			/users/{id} [get]
func (u *userController) GetOneByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := u.service.FindByID(uint(id))
	if err != nil {
		switch err {
		case globalerrors.UserNotFound:
			msg := fmt.Sprintf("User with ID %d not found", id)
			return httpstatus.NotFoundError(msg)
		default:
			return httpstatus.InternalServerError(err.Error())
		}
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

//	 Delete deletes the authenticated user's account.
//	@Summary		Delete logged-in user's account
//	@Description	Delete the account of the authenticated user.
//	@ID				delete-user-by-id
//	@Tags			Users
//	@Security		ApiKeyAuth
//	@Produce		json
//	@Success		200
//	@Success		204
//	@Failure		400	{object}	httpstatus.GlobalErrorHandlerResp	"Invalid id. Should be a positive integer."
//	@Failure		500	{object}	httpstatus.GlobalErrorHandlerResp	"Could not delete user."
//	@Router			/me [delete]
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
		return httpstatus.NoContent("Rows affected: 0")
	}

	return httpstatus.OK("User deleted successfully.")
}

func (u *userController) ExistsOne(c *fiber.Ctx) error {
	username := c.Query("username")
	exists, _ := u.service.ExistsOne(username)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"exists": exists})
}

// UserExists returns true if the user exists by their username or email.
//
//	@Summary		Check if user exists
//	@Description	Check if a user exists by their username or email.
//	@ID				check-user-exists-by-username-or-email
//	@Tags			Users
//	@Produce		json
//	@Param			username	query	string	false	"username"
//	@Param			email		query	string	false	"email"
//	@Router			/users/exists [get]
func (u *userController) UserExists(c *fiber.Ctx) error {
	var exists bool
	username := c.Query("username")
	email := c.Query("email")
	if username == "" && email == "" {
		return httpstatus.BadRequestError("Either username or email must be provided.")
	}

	if username != "" && email != "" {
		return httpstatus.BadRequestError("Only one of username or email must be provided.")
	}

	if username != "" {
		exists, _ = u.service.UsernameExists(username)
	}
	if email != "" {
		exists, _ = u.service.EmailExists(email)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"exists": exists})
}
