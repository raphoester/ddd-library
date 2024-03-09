package driven

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/auth_model"
)

type UsersStorage interface {
	RegisterUser(ctx context.Context, user *auth_model.User) error
}

type TokensStorage interface {
	StoreToken(ctx context.Context, token string) error
}
