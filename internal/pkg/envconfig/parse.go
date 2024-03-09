package envconfig

import (
	"errors"

	"github.com/caarlos0/env/v9"
)

func Parse(in ...any) error {
	var errs error = nil
	for _, v := range in {
		if err := env.Parse(v); err != nil {
			errs = errors.Join(err, errs)
		}
	}
	return errs
}
