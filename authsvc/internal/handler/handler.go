package handler

import (
	"github.com/ikadgzl/inventory/authsvc/internal/config"
	"github.com/ikadgzl/inventory/authsvc/internal/repository"
	"github.com/ikadgzl/inventory/authsvc/internal/util"
	"github.com/ikadgzl/inventory/common/proto/auth"
)

type authHandler struct {
	auth.UnimplementedAuthServiceServer
	jwtUtil util.JwtUtil
	repo    repository.AuthRepository
}

func NewAuthHandler(jwtConfig *config.JWTConfig, repo repository.AuthRepository) auth.AuthServiceServer {
	return &authHandler{
		jwtUtil: util.NewJwtUtil(jwtConfig),
		repo:    repo,
	}
}
