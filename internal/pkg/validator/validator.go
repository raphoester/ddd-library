package validator

import "github.com/go-playground/validator/v10"

func New() *Validator {
	return &Validator{
		Validate: validator.New(),
	}
}

type Validator struct {
	*validator.Validate
}
