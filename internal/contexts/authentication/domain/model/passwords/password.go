package passwords

import (
	"fmt"

	"github.com/raphoester/ddd-library/internal/pkg/cryptoutil"
	"github.com/raphoester/ddd-library/internal/pkg/randomutil"
)

type Password struct {
	hashedPassword string
	salt           string
}

func NewPassword(value string) (*Password, error) {
	var password = &Password{}
	if err := password.GenerateSalt(func() string {
		return randomutil.NewString(10)
	}); err != nil {
		return nil, err
	}

	if err := password.HashAndSet(value); err != nil {
		return nil, err
	}

	return password, nil
}

func (p *Password) GenerateSalt(generator func() string) error {
	// value object is immutable
	if p.salt != "" {
		return fmt.Errorf("salt already exists")
	}

	p.salt = generator()
	return nil
}

func (p *Password) HashAndSet(password string) error {
	// value object is immutable
	if p.hashedPassword != "" {
		return fmt.Errorf("password is already set")
	}

	if len(password) < 8 {
		return fmt.Errorf("password is too short")
	}

	if p.salt == "" {
		return fmt.Errorf("salt is not set")
	}
	password = plaintextPasswordWithSalt(password, p.salt)

	if len(password) > 70 {
		return fmt.Errorf("password is too long")
	}

	p.hashedPassword = cryptoutil.HashPassword(password)

	return nil
}

func (p *Password) Check(password string) (bool, error) {

	passwordAttempt := &Password{salt: p.salt}
	if err := passwordAttempt.HashAndSet(password); err != nil {
		return false, fmt.Errorf("failed to hash password attempt: %w", err)
	}

	return p.hashedPassword == passwordAttempt.hashedPassword, nil
}

func plaintextPasswordWithSalt(plaintextPassword, salt string) string {
	return plaintextPassword + salt
}

func (p *Password) Validate() error {
	if p.hashedPassword == "" {
		return fmt.Errorf("password is empty")
	}

	if p.salt == "" {
		return fmt.Errorf("salt is empty")
	}

	return nil
}
