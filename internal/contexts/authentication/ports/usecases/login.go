package usecases

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/auth_model"
)

type LoginParams struct {
	Email         auth_model.EmailAddress
	PlainPassword string
}

type UsersLoginManager interface {
	Login(ctx context.Context, params LoginParams) (*auth_model.Token, error)
}
