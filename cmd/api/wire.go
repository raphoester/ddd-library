//go:build wireinject

package main

import (
	"github.com/raphoester/ddd-library/internal/contexts/authentication/adapters/inmemory_users_storage"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/controllers/registration_controllers"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/driven"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/usecases/registrations"

	"github.com/google/wire"
)

func getUsersAuthController() *registration_controllers.UsersController {

	wire.Build(
		wire.NewSet(inmemory_users_storage.New,
			registrations.NewUsersRegistrar,
			wire.Bind(
				new(driven.UsersStorage),
				new(*inmemory_users_storage.Repository),
			),

			wire.Bind(
				new(usecases.UsersRegistrar),
				new(*registrations.UsersRegistrar),
			),

			registration_controllers.New,
		),
	)

	return &registration_controllers.UsersController{}
}
