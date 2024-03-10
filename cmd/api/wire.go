//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/tokens"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	usecases2 "github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/usecases/login"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/usecases/registrations"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/adapters/inmemory_tokens_storage"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/adapters/inmemory_users_storage"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/controller"
)

func getUsersAuthController() *controller.Controller {

	wire.Build(
		wire.NewSet(

			inmemory_users_storage.New,
			inmemory_tokens_storage.New,
			registrations.NewUsersRegistrar,
			login.NewUsersLoginManager,

			// repositories
			wire.Bind(
				new(tokens.Storage),
				new(*inmemory_tokens_storage.Repository),
			),

			wire.Bind(
				new(users.Storage),
				new(*inmemory_users_storage.Repository),
			),

			// usecases
			wire.Bind(
				new(usecases2.UsersRegistrar),
				new(*registrations.UsersRegistrar),
			),

			wire.Bind(
				new(usecases2.UsersLoginManager),
				new(*login.UsersLoginManager),
			),

			controller.New,
		),
	)

	return &controller.Controller{}
}
