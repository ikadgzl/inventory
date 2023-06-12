package handler

import (
	"context"

	"github.com/ikadgzl/inventory/common/proto/auth"
)

func (a *authServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	panic("not implemented")
}
