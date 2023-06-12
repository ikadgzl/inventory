package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/common/proto/user"
)

func (h *userHandler) ListUsers(ctx *gin.Context) {
	r, err := h.userSvc.ListUsers(ctx, &user.ListUsersRequest{})
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, r)
}
