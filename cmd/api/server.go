package main

import (
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/raphoester/ddd-library/internal/pkg/grpcutils"
	"github.com/raphoester/ddd-library/internal/pkg/validator"
	"google.golang.org/grpc"
)

const GrpcBufSize = 20 * 1024 * 1024

func getGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				//grpcZap.UnaryServerInterceptor(logger, grpcZap.WithMessageProducer(messageProducer)), // re enable for access logging
				grpcutils.NewRequestValidator(validator.New()),
			),
		),
		grpc.MaxRecvMsgSize(GrpcBufSize),
		grpc.MaxSendMsgSize(GrpcBufSize),
	)

	return grpcServer
}
