package usecases

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/tokens"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
)

type LoginParams struct {
	Email         users.EmailAddress
	PlainPassword string
}

type UsersLoginManager interface {
	Authenticate(ctx context.Context, params LoginParams) (*tokens.Token, error)
}
