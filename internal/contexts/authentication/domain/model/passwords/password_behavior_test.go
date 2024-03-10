package passwords_test

import (
	"testing"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/passwords"
	"github.com/stretchr/testify/require"
)

func TestPassword_CheckPassword(t *testing.T) {
	p := passwords.Password{}

	err := p.GenerateSalt(func() string {
		return "sample-salt"
	})
	require.NoError(t, err)

	pw := "12345678"
	err = p.HashAndSet(pw)
	require.NoError(t, err)

	ok, err := p.Check(pw)
	require.NoError(t, err)
	require.True(t, ok)
}
