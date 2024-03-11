package users

import (
	"fmt"

	"github.com/raphoester/ddd-library/internal/pkg/validator"
)

type EmailAddress string

func NewEmailAddress(value string) (EmailAddress, error) {
	err := validator.New().Var(value, "required,email")
	if err != nil {
		return "", fmt.Errorf("invalid email address: %w", err)
	}

	return EmailAddress(value), nil
}

func (e EmailAddress) String() string {
	return string(e)
}

func (e EmailAddress) validate() error {
	_, err := NewEmailAddress(string(e))
	return err
}
