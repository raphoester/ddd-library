package usecases

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
)

type RegisterUserParams struct {
	Email    users.EmailAddress
	Password users.Password
}

type UsersRegistrar interface {
	RegisterUser(ctx context.Context, params RegisterUserParams) error
}
