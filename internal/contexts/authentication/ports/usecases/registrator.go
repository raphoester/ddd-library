package usecases

import "context"

type RegisterParams struct {
	Email    string
	Password string
}

type UsersRegistrar interface {
	RegisterUser(ctx context.Context, params RegisterParams) error
}
