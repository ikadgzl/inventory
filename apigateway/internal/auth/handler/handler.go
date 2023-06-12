package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/common/proto/auth"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Validate(ctx *gin.Context)
}

type authHandler struct {
	authSvc auth.AuthServiceClient
}

func NewAuthHandler(authSvc auth.AuthServiceClient) AuthHandler {
	return &authHandler{
		authSvc: authSvc,
	}
}
