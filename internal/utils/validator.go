package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct menerima struct dan mengembalikan error human readable
func ValidateStruct(payload interface{}) map[string]string {
	err := validate.Struct(payload)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		switch e.Tag() {
		case "required":
			errors[field] = fmt.Sprintf("%s is required", field)
		case "email":
			errors[field] = fmt.Sprintf("%s must be a valid email", field)
		case "min":
			errors[field] = fmt.Sprintf("%s must be at least %s characters", field, e.Param())
		case "max":
			errors[field] = fmt.Sprintf("%s must be at most %s characters", field, e.Param())
		default:
			errors[field] = fmt.Sprintf("%s is invalid", field)
		}
	}
	return errors
}