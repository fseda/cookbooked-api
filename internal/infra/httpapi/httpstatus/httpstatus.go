package httpstatus

import (
	"github.com/gofiber/fiber/v2"
)

type HTTPStatus interface {
	error
	StatusCode() int
}

type CustomStatus struct {
	Code int
	Message string
}

func (e CustomStatus) Error() string {
	return e.Message
}

func (e CustomStatus) StatusCode() int {
	return e.Code
}

func OK(msg string) CustomStatus {
	return CustomStatus{Code: fiber.StatusOK, Message: msg}
}

func Created(msg string) CustomStatus {
	return CustomStatus{Code: fiber.StatusCreated, Message: msg}
}

func NotFoundError(msg string) CustomStatus {
	return CustomStatus{Code: fiber.StatusNotFound, Message: msg}
}

func BadRequestError(msg string) CustomStatus {
	return CustomStatus{Code: fiber.StatusBadRequest, Message: msg}
}

func UnprocessableEntityError(msg string) CustomStatus {
	return CustomStatus{Code: fiber.StatusUnprocessableEntity, Message: msg}
}

func InternalServerError(msg string) CustomStatus {
	return CustomStatus{Code: fiber.StatusInternalServerError, Message: msg}
}

func ConflictError(msg string) CustomStatus {
	return CustomStatus{Code: fiber.StatusConflict, Message: msg}
}

func UnauthorizedError(msg string) CustomStatus {
	return CustomStatus{Code: fiber.StatusUnauthorized, Message: msg}
}

func ForbiddenError(msg string) CustomStatus {
	return CustomStatus{Code: fiber.StatusForbidden, Message: msg}
}

func MethodNotAllowedError(msg string) CustomStatus {
	return CustomStatus{Code: fiber.StatusMethodNotAllowed, Message: msg}
}

func NoContent(msg string) CustomStatus {
	return CustomStatus{Code: fiber.StatusNoContent, Message: msg}
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Error message"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"

	if e, ok := err.(HTTPStatus); ok {
		code = e.StatusCode()
		msg = e.Error()
	}

	return c.Status(code).JSON(GlobalErrorHandlerResp{
		Success: false,
		Message: msg,
	})
}
