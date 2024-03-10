package controller

import (
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/usecases/login"
)

type Controller struct {
	proto.UnimplementedAuthenticationServer
	usersRegistrar usecases.UsersRegistrar
	loginManager   usecases.UsersLoginManager
}

func New(usersRegistrar usecases.UsersRegistrar, loginManager *login.UsersLoginManager) *Controller {
	return &Controller{
		usersRegistrar: usersRegistrar,
		loginManager:   loginManager,
	}
}
