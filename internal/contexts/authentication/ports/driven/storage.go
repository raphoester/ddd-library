package driven

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/auth_model"
)

type UsersStorage interface {
	RegisterUser(ctx context.Context, user *auth_model.User) error
	FindUserByID(ctx context.Context, id auth_model.ID) (*auth_model.User, error)
	FindUserFromEmail(ctx context.Context, email auth_model.EmailAddress) (*auth_model.User, error)
}

type TokensStorage interface {
	StoreToken(ctx context.Context, token string) error
}
