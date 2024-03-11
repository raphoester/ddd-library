package tokens

import (
	"context"
	"fmt"
	"time"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/id"
)

type Storage interface {
	SaveToken(ctx context.Context, token *TokenDAO) error
	GetFromAccessToken(ctx context.Context, accessToken string) (*TokenDAO, error)
}

func SaveToken(ctx context.Context, storage Storage, token *Token) error {
	return storage.SaveToken(ctx, token.toDAO())
}

func GetFromAccessToken(ctx context.Context, storage Storage, accessToken AccessToken) (*Token, error) {
	dao, err := storage.GetFromAccessToken(ctx, accessToken.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get token from access token: %w", err)
	}

	token, err := dao.toToken()
	if err != nil {
		return nil, fmt.Errorf("failed to convert token dao to token: %w", err)
	}

	return token, nil
}

type TokenDAO struct {
	UserID       id.ID
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

func (t *TokenDAO) toToken() (*Token, error) {
	token := &Token{
		userID:       t.UserID,
		accessToken:  AccessToken(t.AccessToken),
		refreshToken: RefreshToken(t.RefreshToken),
		expiresAt:    TokenExpiration(t.ExpiresAt),
	}

	if err := token.validate(); err != nil {
		return nil, err
	}

	return token, nil
}

func (t *Token) toDAO() *TokenDAO {
	return &TokenDAO{
		UserID:       t.userID,
		AccessToken:  t.accessToken.String(),
		RefreshToken: t.refreshToken.String(),
		ExpiresAt:    t.expiresAt.Time(),
	}
}
