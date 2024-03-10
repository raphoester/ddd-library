package login

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/auth_model"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/driven"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/usecases"
)

func NewUsersLoginManager(
	usersStorage driven.UsersStorage,
	tokensStorage driven.TokensStorage,
) *UsersLoginManager {
	return &UsersLoginManager{
		usersStorage:  usersStorage,
		tokensStorage: tokensStorage,
	}
}

type UsersLoginManager struct {
	usersStorage  driven.UsersStorage
	tokensStorage driven.TokensStorage
}

func (u *UsersLoginManager) Login(ctx context.Context, params usecases.LoginParams) (*auth_model.Token, error) {

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

	token, err := auth_model.NewToken(
		auth_model.NewTokenParams{
			AccessToken:  auth_model.NewAccessToken(),
			RefreshToken: auth_model.NewRefreshToken(),
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
