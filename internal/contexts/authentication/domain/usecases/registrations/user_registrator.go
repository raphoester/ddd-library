package registrations

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
)

type UsersRegistrar struct {
	usersStorage users.Storage
}

func NewUsersRegistrar(usersStorage users.Storage) *UsersRegistrar {
	return &UsersRegistrar{
		usersStorage: usersStorage,
	}
}

func (r *UsersRegistrar) RegisterUser(ctx context.Context, params usecases.RegisterUserParams) error {

	newUserParams := users.NewUserParams{
		EmailAddress: params.Email,
		Password:     params.Password,
		Role:         users.RoleUser,
	}

	user, err := users.NewUser(newUserParams)
	if err != nil {
		return fmt.Errorf("failed to create a new user: %w", err)
	}

	if err := users.Register(ctx, r.usersStorage, user); err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}
