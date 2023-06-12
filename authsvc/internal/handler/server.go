package handler

import (
	"github.com/ikadgzl/inventory/authsvc/internal/config"
	"github.com/ikadgzl/inventory/authsvc/internal/util"
	"github.com/ikadgzl/inventory/common/proto/auth"
)

type authServer struct {
	auth.UnimplementedAuthServiceServer
	jwtUtil util.JwtUtil
}

func NewAuthServer(jwtConfig *config.JWTConfig) auth.AuthServiceServer {
	return &authServer{
		jwtUtil: util.NewJwtUtil(jwtConfig),
	}
}
