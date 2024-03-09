package registrations

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/auth_model"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/driven"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/usecases"
)

type UsersRegistrar struct {
	usersStorage driven.UsersStorage
}

func NewUsersRegistrar(usersStorage driven.UsersStorage) *UsersRegistrar {
	return &UsersRegistrar{
		usersStorage: usersStorage,
	}
}

func (r *UsersRegistrar) RegisterUser(ctx context.Context, params usecases.RegisterParams) error {

	emailAddress, err := auth_model.NewEmailAddress(params.Email)
	if err != nil {
		return fmt.Errorf("failed to create email address: %w", err)
	}

	password, err := auth_model.NewPassword(params.Password)
	if err != nil {
		return fmt.Errorf("failed to create password: %w", err)
	}

	createUserParams := auth_model.NewUserParams{
		EmailAddress: emailAddress,
		Password:     *password,
		Role:         auth_model.RoleUser,
	}

	user, err := auth_model.NewUser(createUserParams)
	if err != nil {
		return fmt.Errorf("failed to create a new user: %w", err)
	}

	if err := r.usersStorage.RegisterUser(ctx, user); err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}
