package users

import (
	"context"
	"fmt"
	"time"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/id"
)

type Storage interface {
	Register(ctx context.Context, user *UserDAO) error
	Get(ctx context.Context, id id.ID) (*UserDAO, error)
	GetFromEmail(ctx context.Context, email string) (*UserDAO, error)
}

func Register(ctx context.Context, storage Storage, user *User) error {
	return storage.Register(ctx, user.toDAO())
}

func Get(ctx context.Context, storage Storage, id id.ID) (*User, error) {
	userDAO, err := storage.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	user, err := userDAO.toUser()
	if err != nil {
		return nil, fmt.Errorf("failed to convert user dao to user: %w", err)
	}

	return user, nil
}

func GetFromEmail(ctx context.Context, storage Storage, email EmailAddress) (*User, error) {
	userDAO, err := storage.GetFromEmail(ctx, email.String())
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	user, err := userDAO.toUser()
	if err != nil {
		return nil, fmt.Errorf("failed to convert user dao to user: %w", err)
	}

	return user, nil
}

type UserDAO struct {
	ID             id.ID
	CreatedAt      time.Time
	IsActive       bool
	Role           string
	EmailAddress   string
	HashedPassword string
	Salt           string
}

func (u *UserDAO) toUser() (*User, error) {
	user := &User{
		id:           u.ID,
		createdAt:    u.CreatedAt,
		isActive:     u.IsActive,
		role:         Role(u.Role),
		emailAddress: EmailAddress(u.EmailAddress),
		password:     Password{hashedPassword: u.HashedPassword, salt: u.Salt},
	}

	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) toDAO() *UserDAO {
	return &UserDAO{
		ID:             u.id,
		CreatedAt:      u.createdAt,
		IsActive:       u.isActive,
		Role:           u.role.String(),
		EmailAddress:   u.emailAddress.String(),
		HashedPassword: u.password.hashedPassword,
		Salt:           u.password.salt,
	}
}
