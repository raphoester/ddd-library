package auth_model

import (
	"errors"
	"fmt"
	"time"

	"github.com/raphoester/ddd-library/internal/pkg/randomutil"
)

type Token struct {
	accessToken  AccessToken
	refreshToken RefreshToken
	expiresAt    TokenExpiration
	userID       ID
}

type AccessToken string

func NewAccessToken() AccessToken {
	return AccessToken(randomutil.NewString(64))
}

func (a AccessToken) Validate() error {
	if len(a) != 64 {
		return errors.New("invalid length")
	}

	return nil
}

type RefreshToken string

func NewRefreshToken() RefreshToken {
	return RefreshToken(randomutil.NewString(64))
}

func (r RefreshToken) Validate() error {
	if len(r) != 64 {
		return errors.New("invalid length")
	}

	return nil
}

type TokenExpiration time.Time

func (t TokenExpiration) Validate() error {
	if time.Time(t).IsZero() {
		return errors.New("expiration time is not set")
	}

	return nil
}

func (t TokenExpiration) IsExpired(now time.Time) bool {
	return time.Time(t).Before(now)
}

type NewTokenParams struct {
	AccessToken  AccessToken
	RefreshToken RefreshToken
	ExpiresAt    TokenExpiration
	ForUserID    ID
}

func (p *NewTokenParams) IsValid() error {
	if err := p.AccessToken.Validate(); err != nil {
		return fmt.Errorf("invalid access token: %w", err)
	}

	if err := p.RefreshToken.Validate(); err != nil {
		return fmt.Errorf("refresh token is invalid: %w", err)
	}

	if err := p.ExpiresAt.Validate(); err != nil {
		return fmt.Errorf("invalid expiration time: %w", err)
	}

	if err := p.ForUserID.Validate(); err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}

	return nil
}

func NewToken(params NewTokenParams) (*Token, error) {
	if err := params.IsValid(); err != nil {
		return nil, fmt.Errorf("invalid access token params: %w", err)
	}

	return &Token{
		accessToken:  params.AccessToken,
		refreshToken: params.RefreshToken,
		expiresAt:    params.ExpiresAt,
		userID:       params.ForUserID,
	}, nil
}

func (a Token) IsExpired(now time.Time) bool {
	return a.expiresAt.IsExpired(now)
}

func (a Token) GetUserID() ID {
	return a.userID
}
