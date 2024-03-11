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

func AccessTokenFromString(str string) (*AccessToken, error) {
	value := AccessToken(str)
	if err := value.validate(); err != nil {
		return nil, err
	}

	return &value, nil
}

func (a AccessToken) validate() error {
	if len(a) != 64 {
		return errors.New("invalid length")
	}

	return nil
}
