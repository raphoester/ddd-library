package usecases

import "context"

type RegisterUserParams struct {
	Email    string
	Password string
}

type UsersRegistrar interface {
	RegisterUser(ctx context.Context, params RegisterUserParams) error
}
