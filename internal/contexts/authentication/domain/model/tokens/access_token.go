package tokens

import (
	"errors"

	"github.com/raphoester/ddd-library/internal/pkg/random"
)

type AccessToken string

func (a AccessToken) String() string {
	return string(a)
}

func NewAccessToken() AccessToken {
	return AccessToken(random.NewString(64))
}

func (a AccessToken) validate() error {
	if len(a) != 64 {
		return errors.New("invalid length")
	}

	return nil
}
