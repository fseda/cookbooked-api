package validation

import (
	"encoding/json"
	"fmt"
	"regexp"
)

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

func (v *Validation) AddValidation(newV Validation) {
	for field, errs := range newV.Errors {
		for _, err := range errs {
			v.AddError(field, fmt.Errorf(err))
		}
	}
}

func (v Validation) HasErrors() bool {
	return len(v.Errors) > 0
}

func (v Validation) IsValid() bool {
	return !v.HasErrors()
}

func (v Validation) String() string {
	return fmt.Sprintf("%v", v.Errors)
}

func (v Validation) ToJsonFormatted() string {
	_4spaces := "    "
	b, _ := json.MarshalIndent(v, "", _4spaces)
	return string(b)
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsEmailLike(input string) bool {
	return emailRegex.MatchString(input)
}
