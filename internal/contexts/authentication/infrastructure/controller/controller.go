package controller

import (
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/usecases/authenticator"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
)

type Controller struct {
	proto.UnimplementedAuthenticationServer

	registrar      usecases.Registrar
	authenticator  usecases.Authenticator
	tokenValidator usecases.TokenValidator
}

func New(usersRegistrar usecases.Registrar, loginManager *authenticator.UsersAuthenticator) *Controller {
	return &Controller{
		registrar:     usersRegistrar,
		authenticator: loginManager,
	}
}
