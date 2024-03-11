package postgresql_users_storage

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/id"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
)

func (r *Repository) Find(ctx context.Context, id id.ID) (*users.UserDAO, error) {
	panic("not implemented")
}

func (r *Repository) FindFromEmail(ctx context.Context, email users.EmailAddress) (*users.User, error) {
	panic("not implemented")
}
