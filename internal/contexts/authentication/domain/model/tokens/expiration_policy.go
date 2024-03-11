package tokens

import (
	"time"

	"github.com/raphoester/ddd-library/internal/pkg/timeutil"
)

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
