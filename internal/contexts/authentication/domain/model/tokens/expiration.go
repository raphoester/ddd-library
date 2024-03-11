package tokens

import (
	"errors"
	"time"
)

type TokenExpiration time.Time

func (t TokenExpiration) validate() error {
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
