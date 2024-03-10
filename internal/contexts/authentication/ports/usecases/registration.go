package usecases

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/auth_model"
)

type RegisterUserParams struct {
	Email    auth_model.EmailAddress
	Password auth_model.Password
}

type UsersRegistrar interface {
	RegisterUser(ctx context.Context, params RegisterUserParams) error
}
