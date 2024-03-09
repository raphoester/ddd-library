package main

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
)

func main() {
	auth := getUsersAuthController()

	_, err := auth.RegisterUser(context.Background(), &proto.RegisterUserRequest{
		Email:    "raphaeloester@gmail.com",
		Password: "12345678",
	})
	if err != nil {
		panic(err)
	}
}
