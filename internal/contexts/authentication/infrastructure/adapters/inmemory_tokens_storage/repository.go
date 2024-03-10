package inmemory_tokens_storage

import (
	"context"
	"sync"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/tokens"
)

type Repository struct {
	tokens map[tokens.AccessToken]*tokens.Token
	mu     sync.Mutex
}

func New() *Repository {
	return &Repository{
		tokens: make(map[tokens.AccessToken]*tokens.Token),
	}
}

func (r *Repository) SaveToken(_ context.Context, token *tokens.Token) error {

	r.mu.Lock()
	r.tokens[token.GetAccessToken()] = token
	r.mu.Unlock()

	return nil
}
