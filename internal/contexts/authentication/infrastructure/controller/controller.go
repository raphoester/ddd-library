package controller

import (
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/usecases/login"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
)

type Controller struct {
	proto.UnimplementedAuthenticationServer
	usersRegistrar usecases.UsersRegistrar
	loginManager   usecases.UsersLoginManager
}

func New(usersRegistrar usecases.UsersRegistrar, loginManager *login.UsersAuthenticator) *Controller {
	return &Controller{
		usersRegistrar: usersRegistrar,
		loginManager:   loginManager,
	}
}
