package controller

import (
	"context"
	"fmt"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/tokens"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
)

func (c *Controller) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	accessToken, err := tokens.AccessTokenFromString(req.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %w", err)
	}

	response, err := c.tokenValidator.ValidateToken(ctx, *accessToken)
	if err != nil {
		return nil, err
	}

	return &proto.ValidateTokenResponse{
		Valid: response.IsValid,
		Role:  response.Role.String(),
	}, nil
}
