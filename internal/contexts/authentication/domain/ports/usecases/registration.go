package usecases

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
)

type RegisterParams struct {
	Email    users.EmailAddress
	Password users.Password
}

type Registrar interface {
	Register(ctx context.Context, params RegisterParams) error
}
