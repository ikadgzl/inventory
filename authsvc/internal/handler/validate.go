package handler

import (
	"context"

	"github.com/ikadgzl/inventory/common/proto/auth"
)

func (a *authHandler) Validate(ctx context.Context, req *auth.ValidateRequest) (*auth.ValidateResponse, error) {
	return nil, nil
}
