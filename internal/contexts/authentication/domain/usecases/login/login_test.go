package login_test

import (
	"context"
	"testing"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/usecases/login"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/adapters/inmemory_tokens_storage"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/adapters/inmemory_users_storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getLoginManager(setup func(
	usersStorage *inmemory_users_storage.Repository,
	tokensStorage *inmemory_tokens_storage.Repository),
) *login.UsersAuthenticator {

	usersStorage := inmemory_users_storage.New()
	tokensStorage := inmemory_tokens_storage.New()

	if setup != nil {
		setup(usersStorage, tokensStorage)
	}

	useCase := login.NewUsersAuthenticator(usersStorage, tokensStorage)
	return useCase
}

func TestLogin_UserDoesNotExist(t *testing.T) {
	useCase := getLoginManager(nil)

	params := usecases.LoginParams{
		Email:         "example@example.test",
		PlainPassword: "password",
	}

	_, err := useCase.Authenticate(context.Background(), params)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find user")
}

func TestLogin_InvalidPassword(t *testing.T) {

	email, _ := users.NewEmailAddress("john.doe@gmail.com")
	password, _ := users.NewPassword("password")

	preMadeUser, err := users.CreateUser(
		users.CreateUserParams{
			EmailAddress: email,
			Password:     *password,
			Role:         users.RoleUser,
		},
	)
	require.NoError(t, err)

	useCase := getLoginManager(
		func(
			usersStorage *inmemory_users_storage.Repository,
			tokensStorage *inmemory_tokens_storage.Repository,
		) {
			err := users.Register(context.Background(), usersStorage, preMadeUser)
			require.NoError(t, err)
		},
	)

	params := usecases.LoginParams{
		Email:         email,
		PlainPassword: "invalid_password",
	}

	_, err = useCase.Authenticate(context.Background(), params)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid password")
}
