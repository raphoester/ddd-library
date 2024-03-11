package inmemory_users_storage

import (
	"context"
	"fmt"
	"sync"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/id"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
)

type Repository struct {
	mu    sync.Mutex
	users map[id.ID]*users.UserDAO
}

func New() *Repository {
	return &Repository{
		users: make(map[id.ID]*users.UserDAO),
	}
}

func (r *Repository) Register(ctx context.Context, user *users.UserDAO) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, existingUser := range r.users {
		if existingUser.EmailAddress == user.EmailAddress {
			return fmt.Errorf("email address already in use")
		}
	}

	r.users[user.ID] = user
	return nil
}

func (r *Repository) Get(ctx context.Context, id id.ID) (*users.UserDAO, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %s not found", id)
	}
	return user, nil
}

func (r *Repository) GetFromEmail(_ context.Context, email string) (*users.UserDAO, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.EmailAddress == email {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user with email %s not found", email)
}
