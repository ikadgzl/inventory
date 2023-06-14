package handler

import (
	"context"
	"errors"

	"github.com/ikadgzl/inventory/authsvc/internal/util"
	"github.com/ikadgzl/inventory/common/proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *authHandler) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	err := util.ValidateLoginReqest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	au, err := a.repo.FindOne(req.Email)
	if err != nil {
		if errors.Is(err, util.ErrUserNotFound) {
			return nil, util.ErrGrpcInvalidCredentials
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	ok := util.CheckPasswordHash(au.Password, req.Password)
	if !ok {
		return nil, util.ErrGrpcInvalidCredentials
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
