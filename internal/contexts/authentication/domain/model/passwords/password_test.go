package passwords

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPassword_GenerateSalt(t *testing.T) {
	sampleSalt := "sample-salt"
	cases := []struct {
		name         string
		generateFunc func() string
		baseState    Password
		wantedSalt   string
		wantErr      bool
	}{
		{
			name: "fail because salt already exists",
			baseState: Password{
				salt:           sampleSalt,
				hashedPassword: "",
			},
			generateFunc: func() string { return sampleSalt },
			wantedSalt:   sampleSalt,
			wantErr:      true,
		}, {
			name: "set and succeed",
			baseState: Password{
				salt:           "",
				hashedPassword: "",
			},
			generateFunc: func() string { return sampleSalt },
			wantedSalt:   sampleSalt,
			wantErr:      false,
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				err := tc.baseState.GenerateSalt(tc.generateFunc)
				if tc.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.wantedSalt, tc.baseState.salt)
				}
			},
		)
	}
}

func TestPassword_HashAndSet(t *testing.T) {
	samplePw := strings.Repeat("123", 10)
	sampleSalt := "sample-salt"
	cases := []struct {
		name            string
		baseState       Password
		password        string
		wantErr         bool
		wantErrContains string
	}{
		{
			name: "fail because salt is not set",
			baseState: Password{
				salt:           "",
				hashedPassword: "",
			},
			password:        samplePw,
			wantErr:         true,
			wantErrContains: "salt is not set",
		}, {
			name: "fail because password is already set",
			baseState: Password{
				salt:           "",
				hashedPassword: samplePw,
			},
			password:        samplePw,
			wantErr:         true,
			wantErrContains: "password is already set",
		}, {
			name: "fail because password is too long",
			baseState: Password{
				salt:           sampleSalt,
				hashedPassword: "",
			},
			password:        strings.Repeat("123", 300),
			wantErr:         true,
			wantErrContains: "password is too long",
		}, {
			name: "fail because password is too short",
			baseState: Password{
				salt:           sampleSalt,
				hashedPassword: "",
			},
			password:        "123",
			wantErr:         true,
			wantErrContains: "password is too short",
		}, {
			name: "succeed",
			baseState: Password{
				salt:           sampleSalt,
				hashedPassword: "",
			},
			password: samplePw,
			wantErr:  false,
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.name, func(t *testing.T) {
				err := tc.baseState.HashAndSet(tc.password)
				if tc.wantErr {
					require.Error(t, err)
					assert.Contains(t, err.Error(), tc.wantErrContains)
				} else {
					require.NoError(t, err)
				}
			},
		)
	}
}
