package tokens

import (
	"time"

	"github.com/raphoester/ddd-library/internal/pkg/time_provider"
)

type TokenExpirationPolicy struct {
	timeProvider        time_provider.Provider
	accessTokenLifetime time.Duration
}

func NewTokenExpirationPolicy(provider time_provider.Provider, lifetime time.Duration) TokenExpirationPolicy {
	return TokenExpirationPolicy{
		timeProvider:        provider,
		accessTokenLifetime: lifetime,
	}
}

func (p TokenExpirationPolicy) NewExpirationTime() TokenExpiration {
	return TokenExpiration(p.timeProvider.Now().Add(p.accessTokenLifetime))
}

var accessTokenExpirationPolicy = NewTokenExpirationPolicy(time_provider.NewActualProvider(), 24*time.Hour)

func SetAccessTokenExpirationPolicy(policy TokenExpirationPolicy) {
	accessTokenExpirationPolicy = policy
}
