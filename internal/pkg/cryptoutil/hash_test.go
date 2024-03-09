package cryptoutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	pw := "12345678_sample-salt"
	hashedPw := HashPassword(pw)
	require.NotEqual(t, hashedPw, pw)
	require.NotEmpty(t, hashedPw)

	// make sure the implementation is deterministic (no builtin salt or whatever)
	reHashed := HashPassword(pw)
	require.Equal(t, reHashed, hashedPw)
}
