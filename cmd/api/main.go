package main

import (
	"net"

	"github.com/raphoester/ddd-library/configs"
	"github.com/raphoester/ddd-library/internal/contexts/authentication/infrastructure/proto"
	"github.com/raphoester/ddd-library/internal/pkg/envconfig"
)

func main() {
	var cfg configs.Config
	if err := envconfig.Parse(&cfg); err != nil {
		panic(err)
	}

	auth := getUsersAuthController()

	srv := getGRPCServer()
	proto.RegisterAuthenticationServer(srv, auth)

	listener, err := net.Listen("tcp", cfg.ServerConfig.BindAddress)
	if err != nil {
		panic(err)
	}

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
