package controller

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/ports/usecases"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (c *Controller) Authenticate(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	params, err := mapAuthenticateParams(req)
	if err != nil {
		return nil, err
	}

	token, err := c.authenticator.Authenticate(ctx, *params)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		Token:        token.AccessToken().String(),
		RefreshToken: token.RefreshToken().String(),
		Expiration:   timestamppb.New(token.GetExpiration().Time()),
	}, nil
}

func mapAuthenticateParams(req *proto.LoginRequest) (*usecases.AuthenticateParams, error) {
	emailAddress, err := users.NewEmailAddress(req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email address: %w", err)
	}

	return &usecases.AuthenticateParams{
		Email:         emailAddress,
		PlainPassword: req.Password,
	}, nil
}
