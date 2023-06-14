package handler

import (
	"context"

	"github.com/ikadgzl/inventory/common/proto/auth"
)

func (a *authHandler) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	claims, err := a.jwtUtil.ValidateToken(req.Token)
	if err != nil {
		return nil, err
	}

	return &auth.ValidateResponse{
		UserId: claims.Id,
	}, nil
}
