package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
		Param       string
	}

	XValidator struct {
		validator *validator.Validate
	}
)

var (
	validate    = validator.New()
	MyValidator = &XValidator{
		validator: validate,
	}
)

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true
			elem.Param = err.Param()

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func (v XValidator) CreateErrorResponse(data interface{}) (errMsgs []string) {
	if errs := v.Validate(data); len(errs) > 0 && errs[0].Error {
		errMsgs = make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s: %s'",
				err.FailedField,
				err.Value,
				err.Tag,
				err.Param,
			))
		}
	}

	return
}

type Validation struct {
	Errors map[string][]string `json:"errors"`
}

func NewValidation() Validation {
	return Validation{
		Errors: make(map[string][]string, 0),
	}
}

func (v *Validation) AddError(field string, err error) {
	v.Errors[field] = append(v.Errors[field], err.Error())
}

func (v *Validation) HasErrors() bool {
	return len(v.Errors) > 0
}

func (v *Validation) IsValid() bool {
	return !v.HasErrors()
}

func (v *Validation) String() string {
	return fmt.Sprintf("%v", v.Errors)
}
