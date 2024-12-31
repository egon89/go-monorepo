package utils

import "github.com/go-playground/validator/v10"

func ValidateStruct(input interface{}) error {
	var validate = validator.New()

	return validate.Struct(input)
}
