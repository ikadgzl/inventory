package handler

import (
	"github.com/ikadgzl/inventory/authsvc/internal/config"
	"github.com/ikadgzl/inventory/authsvc/internal/util"
	"github.com/ikadgzl/inventory/common/proto/auth"
)

type authHandler struct {
	auth.UnimplementedAuthServiceServer
	jwtUtil util.JwtUtil
}

func NewAuthHandler(jwtConfig *config.JWTConfig) auth.AuthServiceServer {
	return &authHandler{
		jwtUtil: util.NewJwtUtil(jwtConfig),
	}
}
