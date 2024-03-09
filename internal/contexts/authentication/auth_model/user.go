package auth_model

import (
	"errors"
	"fmt"
	"time"

	"github.com/raphoester/ddd-library/internal/pkg/timeutil"
	"github.com/raphoester/ddd-library/internal/pkg/validator"
)

type User struct {
	id           ID
	role         Role
	emailAddress EmailAddress
	password     Password
	createdAt    time.Time
	isActive     bool
}

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

func NewRole(value string) (Role, error) {
	role := Role(value)
	switch role {
	case RoleAdmin, RoleUser:
		return role, nil
	default:
		return "", errors.New("role does not exist")
	}
}

func (r Role) Validate() error {
	if _, err := NewRole(string(r)); err != nil {
		return err
	}

	return nil
}

type EmailAddress string

func NewEmailAddress(value string) (EmailAddress, error) {
	err := validator.New().Var(value, "required,email")
	if err != nil {
		return "", fmt.Errorf("invalid email address: %w", err)
	}

	return EmailAddress(value), nil
}

func (e EmailAddress) Validate() error {
	_, err := NewEmailAddress(string(e))
	return err
}

type NewUserParams struct {
	EmailAddress EmailAddress
	Role         Role
	Password     Password
}

func (p NewUserParams) Validate() error {
	if err := p.EmailAddress.Validate(); err != nil {
		return fmt.Errorf("invalid email address: %w", err)
	}

	if err := p.Role.Validate(); err != nil {
		return fmt.Errorf("invalid role: %w", err)
	}

	if err := p.Password.Validate(); err != nil {
		return fmt.Errorf("valid password: %w", err)
	}

	return nil
}

func NewUser(params NewUserParams) (*User, error) {

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("invalid user params: %w", err)
	}

	return &User{
		id:           NewID(),
		createdAt:    timeutil.Now(),
		isActive:     false,
		role:         params.Role,
		emailAddress: params.EmailAddress,
		password:     params.Password,
	}, nil
}

func (u *User) CheckPassword(pw string) (bool, error) {
	return u.password.Check(pw)
}

func (u *User) GetID() ID {
	return u.id
}
