package registrations_test

import (
	"context"
	"strings"
	"testing"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/usecases/registrations"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/adapters/inmemory_users_storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsersRegistrar_RegisterUser(t *testing.T) {
	testEmail, _ := users.NewEmailAddress("test@example.com")
	testPassword, _ := users.NewPassword("password")

	cases := []struct {
		name                  string
		params                usecases.RegisterUserParams
		setupFunc             func(*inmemory_users_storage.Repository)
		expectError           bool
		expectedErrorContains string
	}{
		{
			name: "user already exists with email",
			setupFunc: func(usersStorage *inmemory_users_storage.Repository) {
				password, err := users.NewPassword("password")
				require.NoError(t, err)

				user, err := users.NewUser(
					users.NewUserParams{
						Role:         users.RoleUser,
						EmailAddress: testEmail,
						Password:     *password,
					},
				)
				require.NoError(t, err)

				err = users.Register(context.Background(), usersStorage, user)
				require.NoError(t, err)
			},
			params: usecases.RegisterUserParams{
				Email:    testEmail,
				Password: *testPassword,
			},
			expectError:           true,
			expectedErrorContains: "email address already in use",
		}, {
			name: "valid",
			params: usecases.RegisterUserParams{
				Email:    testEmail,
				Password: *testPassword,
			},
			expectError: false,
		},
	}

	for _, c := range cases {
		t.Run(
			c.name, func(t *testing.T) {

				usersStorage := inmemory_users_storage.New()
				if c.setupFunc != nil {
					c.setupFunc(usersStorage)
				}
				registrar := registrations.NewUsersRegistrar(usersStorage)

				err := registrar.RegisterUser(context.Background(), c.params)
				if c.expectError {
					require.Error(t, err)
					assert.True(t, strings.Contains(err.Error(), c.expectedErrorContains))
				} else {
					require.NoError(t, err)
					_, err := users.FindFromEmail(context.Background(), usersStorage, c.params.Email)
					require.NoError(t, err)
				}
			},
		)
	}
}
