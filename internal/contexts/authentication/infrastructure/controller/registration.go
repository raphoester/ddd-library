package controller

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/auth_model"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/usecases"
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
	password, err := auth_model.NewPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create password: %w", err)
	}

	email, err := auth_model.NewEmailAddress(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to create email address: %w", err)
	}

	return &usecases.RegisterUserParams{
		Email:    email,
		Password: *password,
	}, nil
}
