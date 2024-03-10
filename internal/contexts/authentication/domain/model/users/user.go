package users

import (
	"errors"
	"fmt"
	"time"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/id"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/passwords"
	"github.com/raphoester/ddd-library/internal/pkg/timeutil"
)

type User struct {
	id           id.ID
	role         Role
	emailAddress EmailAddress
	password     passwords.Password
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

type NewUserParams struct {
	EmailAddress EmailAddress
	Role         Role
	Password     passwords.Password
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
		id:           id.NewID(),
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

func (u *User) GetID() id.ID {
	return u.id
}

func (u *User) HasEmailAddress(email EmailAddress) bool {
	return u.emailAddress == email
}

func (u *User) GetEmailAddress() EmailAddress {
	return u.emailAddress
}

func (u *User) IsActive() bool {
	return u.isActive
}
