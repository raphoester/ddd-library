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

func (u *User) Validate() error {
	if err := u.id.Validate(); err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	if err := u.role.Validate(); err != nil {
		return fmt.Errorf("invalid role: %w", err)
	}

	if err := u.emailAddress.Validate(); err != nil {
		return fmt.Errorf("invalid email address: %w", err)
	}

	if err := u.password.Validate(); err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}

	if u.createdAt.IsZero() {
		return errors.New("invalid created at time")
	}

	return nil
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

func (e EmailAddress) String() string {
	return string(e)
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

func (u *User) HasEmailAddress(email EmailAddress) bool {
	return u.emailAddress == email
}

func (u *User) GetEmailAddress() EmailAddress {
	return u.emailAddress
}
