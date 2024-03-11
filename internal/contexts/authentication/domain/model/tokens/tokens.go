package tokens

import (
	"fmt"
	"time"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/id"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
)

type Token struct {
	accessToken  AccessToken
	refreshToken RefreshToken
	expiresAt    TokenExpiration
	userID       id.ID
}

func (t *Token) validate() error {
	if err := t.accessToken.validate(); err != nil {
		return fmt.Errorf("invalid access token: %w", err)
	}

	if err := t.refreshToken.validate(); err != nil {
		return fmt.Errorf("invalid refresh token: %w", err)
	}

	if err := t.expiresAt.validate(); err != nil {
		return fmt.Errorf("invalid expiration time: %w", err)
	}

	if err := t.userID.Validate(); err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	return nil
}

func (t *Token) GetAccessToken() AccessToken {
	return t.accessToken
}

type CreateTokenParams struct {
	AccessToken  AccessToken
	RefreshToken RefreshToken
	ForUser      users.User
}

func CreateToken(params CreateTokenParams) (*Token, error) {

	token := &Token{
		accessToken:  params.AccessToken,
		refreshToken: params.RefreshToken,
		expiresAt:    accessTokenExpirationPolicy.NewExpirationTime(),
		userID:       params.ForUser.GetID(),
	}

	if err := token.validate(); err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return token, nil
}

func (t *Token) IsExpired(now time.Time) bool {
	return t.expiresAt.IsExpired(now)
}

func (t *Token) GetRefreshToken() RefreshToken {
	return t.refreshToken
}

func (t *Token) GetExpiration() TokenExpiration {
	return t.expiresAt
}

func (t *Token) GetUserID() id.ID {
	return t.userID
}
