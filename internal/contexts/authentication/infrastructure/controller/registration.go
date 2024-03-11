package controller

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
)

func (c *Controller) RegisterUser(ctx context.Context,
	req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {

	params, err := mapRegisterUserParams(req)
	if err != nil {
		return nil, err
	}

	if err := c.usersRegistrar.RegisterUser(ctx, *params); err != nil {
		return nil, err
	}

	return &proto.RegisterUserResponse{}, nil
}

func mapRegisterUserParams(req *proto.RegisterUserRequest) (*usecases.RegisterUserParams, error) {
	password, err := users.NewPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create password: %w", err)
	}

	email, err := users.NewEmailAddress(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to create email address: %w", err)
	}

	return &usecases.RegisterUserParams{
		Email:    email,
		Password: *password,
	}, nil
}
