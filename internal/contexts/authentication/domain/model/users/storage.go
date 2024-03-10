package users

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/id"
)

type Storage interface {
	RegisterUser(ctx context.Context, user *User) error
	FindUserByID(ctx context.Context, id id.ID) (*User, error)
	FindUserFromEmail(ctx context.Context, email EmailAddress) (*User, error)
}
