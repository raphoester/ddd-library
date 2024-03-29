package usecases

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/tokens"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
)

type AuthenticateParams struct {
	Email         users.EmailAddress
	PlainPassword string
}

type Authenticator interface {
	Authenticate(ctx context.Context, params AuthenticateParams) (*tokens.Token, error)
}
