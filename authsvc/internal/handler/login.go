package handler

import (
	"context"

	"github.com/ikadgzl/inventory/common/proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *authHandler) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	if req.Username == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid credentials")
	}

	if req.Username != "admin" || req.Password != "admin" {
		return nil, status.Error(codes.Unauthenticated, "invalid username or password")
	}

	// TODO: implement login logic
	// 1. check username and password
	// 2. generate token
	// 4. return token to client

	res := &auth.AuthResponse{
		Token: "mytoken",
	}

	return res, nil
}
