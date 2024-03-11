package authenticator

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/tokens"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
)

func New(
	usersStorage users.Storage,
	tokensStorage tokens.Storage,
) *UsersAuthenticator {
	return &UsersAuthenticator{
		usersStorage:  usersStorage,
		tokensStorage: tokensStorage,
	}
}

type UsersAuthenticator struct {
	usersStorage  users.Storage
	tokensStorage tokens.Storage
}

func (u *UsersAuthenticator) Authenticate(ctx context.Context, params usecases.AuthenticateParams) (*tokens.Token, error) {

	user, err := users.GetFromEmail(ctx, u.usersStorage, params.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	ok, err := user.CheckPassword(params.PlainPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to check password: %w", err)
	}

	if !ok {
		return nil, fmt.Errorf("invalid password")
	}

	token, err := tokens.CreateToken(
		tokens.CreateTokenParams{
			AccessToken:  tokens.NewAccessToken(),
			RefreshToken: tokens.NewRefreshToken(),
			ForUser:      *user,
		},
	)
	if err != nil {
		return nil, err
	}

	if err := tokens.SaveToken(ctx, u.tokensStorage, token); err != nil {
		return nil, err
	}

	return token, nil
}
