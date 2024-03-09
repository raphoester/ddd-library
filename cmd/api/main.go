package main

import (
	"context"

	"github.com/raphoester/ddd-library/internal/contexts/authentication/ports/usecases"
)

func main() {
	auth := getAuthUseCase()

	if err := auth.RegisterUser(context.Background(), usecases.RegisterParams{
		Email:    "raphaeloester@gmail.com",
		Password: "123456789101112",
	}); err != nil {
		panic(err)
	}
}
