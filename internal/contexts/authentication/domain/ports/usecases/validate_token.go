package usecases

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/tokens"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
)

type ValidateTokenResponse struct {
	IsValid bool
	Role    users.Role
}

type TokenValidator interface {
	ValidateToken(ctx context.Context, accessToken tokens.AccessToken) (*ValidateTokenResponse, error)
}
