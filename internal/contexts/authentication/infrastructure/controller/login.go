package controller

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *Controller) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	params, err := mapLoginParams(req)
	if err != nil {
		return nil, err
	}

	token, err := c.loginManager.Login(ctx, *params)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		Token:      token.GetAccessToken().String(),
		Expiration: timestamppb.New(token.GetExpiration().Time()),
	}, nil
}

func mapLoginParams(req *proto.LoginRequest) (*usecases.LoginParams, error) {
	emailAddress, err := users.NewEmailAddress(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to create email address: %w", err)
	}

	return &usecases.LoginParams{
		Email:         emailAddress,
		PlainPassword: req.Password,
	}, nil
}
