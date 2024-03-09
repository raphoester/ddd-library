package registration_controllers

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/usecases"
)

type UsersController struct {
	proto.UnimplementedAuthenticationServer
	usersRegistrar usecases.UsersRegistrar
}

func New(usersRegistrar usecases.UsersRegistrar) *UsersController {
	return &UsersController{
		usersRegistrar: usersRegistrar,
	}
}

func (c *UsersController) RegisterUser(ctx context.Context,
	req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {

	if err := c.usersRegistrar.RegisterUser(ctx, usecases.RegisterUserParams{
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		return nil, err
	}

	return &proto.RegisterUserResponse{}, nil
}
