package login

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/tokens"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
)

func NewUsersLoginManager(
	usersStorage users.Storage,
	tokensStorage tokens.Storage,
) *UsersLoginManager {
	return &UsersLoginManager{
		usersStorage:  usersStorage,
		tokensStorage: tokensStorage,
	}
}

type UsersLoginManager struct {
	usersStorage  users.Storage
	tokensStorage tokens.Storage
}

func (u *UsersLoginManager) Login(ctx context.Context, params usecases.LoginParams) (*tokens.Token, error) {

	user, err := u.usersStorage.FindUserFromEmail(ctx, params.Email)
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

	token, err := tokens.NewToken(
		tokens.NewTokenParams{
			AccessToken:  tokens.NewAccessToken(),
			RefreshToken: tokens.NewRefreshToken(),
			ForUser:      *user,
		},
	)
	if err != nil {
		return nil, err
	}

	if err := u.tokensStorage.SaveToken(ctx, token); err != nil {
		return nil, err
	}

	return token, nil
}
