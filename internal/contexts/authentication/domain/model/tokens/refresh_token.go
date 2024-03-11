package tokens

import (
	"errors"

	"github.com/raphoester/ddd-library/internal/pkg/random"
)

type RefreshToken string

func NewRefreshToken() RefreshToken {
	return RefreshToken(random.NewString(64))
}

func (r RefreshToken) validate() error {
	if len(r) != 64 {
		return errors.New("invalid length")
	}

	return nil
}
