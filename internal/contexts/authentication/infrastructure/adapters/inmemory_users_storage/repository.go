package inmemory_users_storage

import (
	"context"
	"fmt"
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
	defer r.mu.Unlock()
	for _, existingUser := range r.users {
		if existingUser.GetEmailAddress() == user.GetEmailAddress() {
			return fmt.Errorf("email address already in use")
		}
	}
	r.users[user.GetID()] = user
	return nil
}

func (r *Repository) FindUserByID(_ context.Context, id auth_model.ID) (*auth_model.User, error) {
	r.mu.Lock()
	user, ok := r.users[id]
	r.mu.Unlock()
	if !ok {
		return nil, fmt.Errorf("user with id %s not found", id)
	}
	return user, nil
}

func (r *Repository) FindUserFromEmail(_ context.Context, email auth_model.EmailAddress) (*auth_model.User, error) {
	r.mu.Lock()
	for _, user := range r.users {
		if user.HasEmailAddress(email) {
			return user, nil
		}
	}
	r.mu.Unlock()

	return nil, fmt.Errorf("user with email %s not found", email)
}
