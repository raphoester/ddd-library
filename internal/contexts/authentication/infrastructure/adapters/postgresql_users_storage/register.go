package postgresql_users_storage

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/domain/model/users"
)

func (r *Repository) Register(ctx context.Context, user *users.UserDAO) error {

	query := `INSERT INTO users (
		id, 
		email_address,
		hashed_password, 
		salt,
		created_at, 
		is_active, 
		role
	) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.EmailAddress,
		user.HashedPassword,
		user.Salt,
		user.CreatedAt,
		user.IsActive,
		user.Role,
	)

	if err != nil {
		return err
	}

	return nil
}
