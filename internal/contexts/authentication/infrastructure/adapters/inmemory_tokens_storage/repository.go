package inmemory_tokens_storage

import (
	"context"
	"sync"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/auth_model"
)

type Repository struct {
	tokens map[auth_model.AccessToken]*auth_model.Token
	mu     sync.Mutex
}

func New() *Repository {
	return &Repository{
		tokens: make(map[auth_model.AccessToken]*auth_model.Token),
	}
}

func (r *Repository) SaveToken(_ context.Context, token *auth_model.Token) error {

	r.mu.Lock()
	r.tokens[token.GetAccessToken()] = token
	r.mu.Unlock()

	return nil
}
