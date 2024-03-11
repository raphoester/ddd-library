//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/tokens"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/usecases/authenticator"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/usecases/registrar"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/usecases/token_validator"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/adapters/inmemory_tokens_storage"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/adapters/inmemory_users_storage"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/controller"
)

func getUsersAuthController() *controller.Controller {

	wire.Build(
		wire.NewSet(

			inmemory_users_storage.New,
			inmemory_tokens_storage.New,
			registrar.New,
			authenticator.New,
			token_validator.New,

			// repositories
			wire.Bind(
				new(tokens.Storage),
				new(*inmemory_tokens_storage.Repository),
			),

			wire.Bind(
				new(users.Storage),
				new(*inmemory_users_storage.Repository),
			),

			// use cases
			wire.Bind(
				new(usecases.Registrar),
				new(*registrar.Registrar),
			),

			wire.Bind(
				new(usecases.Authenticator),
				new(*authenticator.UsersAuthenticator),
			),

			wire.Bind(
				new(usecases.TokenValidator),
				new(*token_validator.TokenValidator),
			),

			controller.New,
		),
	)

	return &controller.Controller{}
}
