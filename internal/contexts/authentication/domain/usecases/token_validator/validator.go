package token_validator

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/tokens"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/pkg/time_provider"
)

type TokenValidator struct {
	tokenStorage tokens.Storage
	usersStorage users.Storage
}

func New(tokenStorage tokens.Storage, usersStorage users.Storage) *TokenValidator {
	return &TokenValidator{
		tokenStorage: tokenStorage,
		usersStorage: usersStorage,
	}
}

func (v *TokenValidator) ValidateToken(ctx context.Context,
	accessToken tokens.AccessToken) (*usecases.ValidateTokenResponse, error) {

	token, err := tokens.GetFromAccessToken(ctx, v.tokenStorage, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	if token.IsExpired(time_provider.Now()) {
		return &usecases.ValidateTokenResponse{
			IsValid: false,
		}, nil
	}

	user, err := users.Get(ctx, v.usersStorage, token.UserID())
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &usecases.ValidateTokenResponse{
		IsValid: true,
		Role:    user.Role(),
	}, nil
}
