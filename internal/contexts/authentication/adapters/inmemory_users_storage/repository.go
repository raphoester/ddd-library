package inmemory_users_storage

import (
	"context"
	"sync"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/auth_model"
)

type Repository struct {
	mu    sync.Mutex
	users map[auth_model.ID]*auth_model.User
}

func New() *Repository {
	return &Repository{
		users: make(map[auth_model.ID]*auth_model.User),
	}
}

func (r *Repository) RegisterUser(_ context.Context, user *auth_model.User) error {
	r.mu.Lock()
	r.users[user.GetID()] = user
	r.mu.Unlock()
	return nil
}
