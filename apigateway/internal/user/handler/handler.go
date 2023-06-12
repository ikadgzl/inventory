package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/common/proto/user"
)

type UserHandler interface {
	ListUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userHandler struct {
	userSvc user.UserServiceClient
}

func NewUserHandler(userSvc user.UserServiceClient) UserHandler {
	return &userHandler{
		userSvc: userSvc,
	}
}
