package grpc_validator

import (
	"context"

	validatorV10 "github.com/go-playground/validator/v10"
	"github.com/raphoester/ddd-library/internal/pkg/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRequestValidator(validator *validator.Validator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if err := validator.Struct(req); err != nil {
			return nil, validatorErrorToGRPCError(err, info.FullMethod)
		}

		return handler(ctx, req)
	}
}

func validatorErrorToGRPCError(err error, structName string) error {
	br := &errdetails.BadRequest{}

	if ve, ok := err.(validatorV10.ValidationErrors); ok {
		for _, fe := range ve {
			br.FieldViolations = append(
				br.FieldViolations, &errdetails.BadRequest_FieldViolation{
					Field:       fe.Field(),
					Description: fe.Error(),
				},
			)
		}
	}

	st := status.New(codes.InvalidArgument, "invalid "+structName)
	st, _ = st.WithDetails(br)
	return st.Err()
}
