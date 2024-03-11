package users

import (
	"errors"
	"fmt"
	"time"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/id"
	"github.com/raphoester/ddd-library/internal/pkg/time_provider"
)

type User struct {
	id           id.ID
	role         Role
	emailAddress EmailAddress
	password     Password
	createdAt    time.Time
	isActive     bool
}

func (u *User) validate() error {
	if err := u.id.Validate(); err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	if err := u.role.validate(); err != nil {
		return fmt.Errorf("invalid role: %w", err)
	}

	if err := u.emailAddress.validate(); err != nil {
		return fmt.Errorf("invalid email address: %w", err)
	}

	if err := u.password.validate(); err != nil {
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
	Password     Password
}

func NewUser(params NewUserParams) (*User, error) {

	user := &User{
		id:           id.Create(),
		createdAt:    time_provider.Now(),
		isActive:     false,
		role:         params.Role,
		emailAddress: params.EmailAddress,
		password:     params.Password,
	}

	if err := user.validate(); err != nil {
		return nil, fmt.Errorf("invalid user: %w", err)
	}

	return user, nil
}

func (u *User) CheckPassword(pw string) (bool, error) {
	return u.password.Check(pw)
}

func (u *User) GetID() id.ID {
	return u.id
}

func (u *User) IsActive() bool {
	return u.isActive
}
