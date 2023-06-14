package handler

import (
	"context"
	"errors"

	"github.com/ikadgzl/inventory/authsvc/internal/util"
	"github.com/ikadgzl/inventory/common/proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *authHandler) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.AuthResponse, error) {
	err := util.ValidateRegisterRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	au, err := a.repo.FindOne(req.Email)
	if err != nil && !errors.Is(err, util.ErrUserNotFound) {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if au != nil {
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = a.repo.Create(req.Name, req.Email, hashedPassword)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	t, err := a.jwtUtil.GenerateToken(au.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &auth.AuthResponse{
		Token: t,
	}

	return res, nil
}
