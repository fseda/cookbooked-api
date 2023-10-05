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

func MethodNotAllowedError(msg string) CustomError {
	return CustomError{Code: fiber.StatusMethodNotAllowed, Message: msg}
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
