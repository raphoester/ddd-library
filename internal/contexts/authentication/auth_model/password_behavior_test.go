package auth_model_test

import (
	"testing"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/auth_model"

	"github.com/stretchr/testify/require"
)

func TestPassword_CheckPassword(t *testing.T) {
	p := auth_model.Password{}

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
