//go:build wireinject

package main

import (
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/adapters/inmemory_tokens_storage"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/adapters/inmemory_users_storage"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/controller"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/driven"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/usecases/login"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/usecases/registrations"

	"github.com/google/wire"
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
				new(driven.TokensStorage),
				new(*inmemory_tokens_storage.Repository),
			),

			wire.Bind(
				new(driven.UsersStorage),
				new(*inmemory_users_storage.Repository),
			),

			// usecases
			wire.Bind(
				new(usecases.UsersRegistrar),
				new(*registrations.UsersRegistrar),
			),

			wire.Bind(
				new(usecases.UsersLoginManager),
				new(*login.UsersLoginManager),
			),

			controller.New,
		),
	)

	return &controller.Controller{}
}
