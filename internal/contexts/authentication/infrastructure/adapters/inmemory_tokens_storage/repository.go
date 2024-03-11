package inmemory_tokens_storage

import (
	"context"
	"errors"
	"sync"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/tokens"
)

type Repository struct {
	tokens map[string]*tokens.TokenDAO
	mu     sync.Mutex
}

func New() *Repository {
	return &Repository{
		tokens: make(map[string]*tokens.TokenDAO),
	}
}

func (r *Repository) SaveToken(_ context.Context, token *tokens.TokenDAO) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.tokens[token.AccessToken] = token

	return nil
}

func (r *Repository) GetFromAccessToken(_ context.Context, accessToken string) (*tokens.TokenDAO, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	token, ok := r.tokens[accessToken]
	if !ok {
		return nil, errors.New("token not found")
	}
	return token, nil
}
