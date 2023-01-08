package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(i interface{}) error {
	validate := validator.New()
	err := validate.Struct(i)

	return err
}

func IsValidationError(err error) bool {
	_, ok := err.(validator.ValidationErrors)

	return ok
}

func ValidatorErrors(err error) *map[string]string {
	// Define variable for error fields.
	errFields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		// Append error message to the map, where key is a field name,
		// and value is an error description.
		var errMessage string
		switch err.Tag() {
		case "required":
			errMessage = "is required"
		case "min":
			errMessage = fmt.Sprintf("must have a minimum value of %s", err.Param())
		case "max":
			errMessage = fmt.Sprintf("must have a maximum value of %s", err.Param())
		case "email":
			errMessage = "must be a valid email address"
		default:
			errMessage = fmt.Sprintf("must match the '%s' validation rule", err.Tag())
		}

		errFields[err.Field()] = errMessage
	}

	return &errFields
}
