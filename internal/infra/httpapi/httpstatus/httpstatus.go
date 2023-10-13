package httpstatus

import (
	"github.com/gofiber/fiber/v2"
)

type HTTPError interface {
	error
	StatusCode() int
}

type CustomError struct {
	Code int
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}

func (e CustomError) StatusCode() int {
	return e.Code
}

func NotFoundError(msg string) CustomError {
	return CustomError{Code: fiber.StatusNotFound, Message: msg}
}

func BadRequestError(msg string) CustomError {
	return CustomError{Code: fiber.StatusBadRequest, Message: msg}
}

func UnprocessableEntityError(msg string) CustomError {
	return CustomError{Code: fiber.StatusUnprocessableEntity, Message: msg}
}

func InternalServerError(msg string) CustomError {
	return CustomError{Code: fiber.StatusInternalServerError, Message: msg}
}

func ConflictError(msg string) CustomError {
	return CustomError{Code: fiber.StatusConflict, Message: msg}
}

func UnauthorizedError(msg string) CustomError {
	return CustomError{Code: fiber.StatusUnauthorized, Message: msg}
}

func ForbiddenError(msg string) CustomError {
	return CustomError{Code: fiber.StatusForbidden, Message: msg}
}

func MethodNotAllowedError(msg string) CustomError {
	return CustomError{Code: fiber.StatusMethodNotAllowed, Message: msg}
}

func NoContent(msg string) CustomError {
	return CustomError{Code: fiber.StatusNoContent, Message: msg}
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"

	if e, ok := err.(HTTPError); ok {
		code = e.StatusCode()
		msg = e.Error()
	}

	return c.Status(code).JSON(GlobalErrorHandlerResp{
		Success: false,
		Message: msg,
	})
}
