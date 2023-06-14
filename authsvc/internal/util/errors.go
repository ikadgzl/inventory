package util

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrGrpcInvalidCredentials = status.Error(codes.InvalidArgument, "invalid credentials")
)
