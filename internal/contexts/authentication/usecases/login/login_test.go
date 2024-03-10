package login_test

import (
	"context"
	"testing"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/auth_model"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/adapters/inmemory_tokens_storage"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/adapters/inmemory_users_storage"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/usecases/login"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getLoginManager(setup func(
	usersStorage *inmemory_users_storage.Repository,
	tokensStorage *inmemory_tokens_storage.Repository),
) *login.UsersLoginManager {

	usersStorage := inmemory_users_storage.New()
	tokensStorage := inmemory_tokens_storage.New()

	if setup != nil {
		setup(usersStorage, tokensStorage)
	}

	useCase := login.NewUsersLoginManager(usersStorage, tokensStorage)
	return useCase
}

func TestLogin_UserDoesNotExist(t *testing.T) {
	useCase := getLoginManager(nil)

	params := usecases.LoginParams{
		Email:         "example@example.test",
		PlainPassword: "password",
	}

	_, err := useCase.Login(context.Background(), params)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find user")
}

func TestLogin_InvalidPassword(t *testing.T) {

	email, _ := auth_model.NewEmailAddress("john.doe@gmail.com")
	password, _ := auth_model.NewPassword("password")

	preMadeUser, err := auth_model.NewUser(
		auth_model.NewUserParams{
			EmailAddress: email,
			Password:     *password,
			Role:         auth_model.RoleUser,
		},
	)
	require.NoError(t, err)

	useCase := getLoginManager(
		func(
			usersStorage *inmemory_users_storage.Repository,
			tokensStorage *inmemory_tokens_storage.Repository,
		) {
			err := usersStorage.RegisterUser(context.Background(), preMadeUser)
			require.NoError(t, err)
		},
	)

	params := usecases.LoginParams{
		Email:         email,
		PlainPassword: "invalid_password",
	}

	_, err = useCase.Login(context.Background(), params)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid password")
}
