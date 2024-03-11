package controller

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
)

func (c *Controller) Register(ctx context.Context,
	req *proto.RegisterRequest) (*proto.RegisterResponse, error) {

	params, err := mapRegisterParams(req)
	if err != nil {
		return nil, err
	}

	if err := c.registrar.Register(ctx, *params); err != nil {
		return nil, err
	}

	return &proto.RegisterResponse{}, nil
}

func mapRegisterParams(req *proto.RegisterRequest) (*usecases.RegisterParams, error) {
	password, err := users.NewPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create password: %w", err)
	}

	email, err := users.NewEmailAddress(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to create email address: %w", err)
	}

	return &usecases.RegisterParams{
		Email:    email,
		Password: *password,
	}, nil
}
