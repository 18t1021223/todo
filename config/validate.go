package config

import "github.com/go-playground/validator/v10"

var ValidatorG *validator.Validate

func InitValidation() {
	ValidatorG = validator.New(validator.WithRequiredStructEnabled())
}
