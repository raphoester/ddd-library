package auth_model

import (
	"errors"
	"fmt"
	"time"

	"github.com/raphoester/ddd-library/internal/pkg/randomutil"
	"github.com/raphoester/ddd-library/internal/pkg/timeutil"
)

type Token struct {
	accessToken  AccessToken
	refreshToken RefreshToken
	expiresAt    TokenExpiration
	userID       ID
}

func (t *Token) GetAccessToken() AccessToken {
	return t.accessToken
}

type AccessToken string

func (a AccessToken) String() string {
	return string(a)
}

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

func (t TokenExpiration) Time() time.Time {
	return time.Time(t)
}

func (t TokenExpiration) IsExpired(now time.Time) bool {
	return time.Time(t).Before(now)
}

type NewTokenParams struct {
	AccessToken  AccessToken
	RefreshToken RefreshToken
	ForUser      User
}

func (p *NewTokenParams) IsValid() error {
	if err := p.AccessToken.Validate(); err != nil {
		return fmt.Errorf("invalid access token: %w", err)
	}

	if err := p.RefreshToken.Validate(); err != nil {
		return fmt.Errorf("refresh token is invalid: %w", err)
	}

	if err := p.ForUser.Validate(); err != nil {
		return fmt.Errorf("invalid user: %w", err)
	}

	if !p.ForUser.isActive {
		return errors.New("user is not active")
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
		expiresAt:    accessTokenExpirationPolicy.NewExpirationTime(),
		userID:       params.ForUser.id,
	}, nil
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

func (t *Token) GetUserID() ID {
	return t.userID
}

type TokenExpirationPolicy struct {
	timeProvider        timeutil.Provider
	accessTokenLifetime time.Duration
}

func NewTokenExpirationPolicy(provider timeutil.Provider, lifetime time.Duration) TokenExpirationPolicy {
	return TokenExpirationPolicy{
		timeProvider:        provider,
		accessTokenLifetime: lifetime,
	}
}

func (p TokenExpirationPolicy) NewExpirationTime() TokenExpiration {
	return TokenExpiration(p.timeProvider.Now().Add(p.accessTokenLifetime))
}

var accessTokenExpirationPolicy = NewTokenExpirationPolicy(timeutil.NewActualProvider(), 24*time.Hour)

func SetAccessTokenExpirationPolicy(policy TokenExpirationPolicy) {
	accessTokenExpirationPolicy = policy
}
