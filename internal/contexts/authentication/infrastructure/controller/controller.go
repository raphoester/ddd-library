package controller

import (
	usecases2 "github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/usecases/login"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
)

type Controller struct {
	proto.UnimplementedAuthenticationServer
	usersRegistrar usecases2.UsersRegistrar
	loginManager   usecases2.UsersLoginManager
}

func New(usersRegistrar usecases2.UsersRegistrar, loginManager *login.UsersLoginManager) *Controller {
	return &Controller{
		usersRegistrar: usersRegistrar,
		loginManager:   loginManager,
	}
}
