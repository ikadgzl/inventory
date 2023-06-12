package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/common/proto/user"
)

func (h *userHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	r, err := h.userSvc.DeleteUser(ctx, &user.DeleteUserRequest{
		Id: id,
	})
	if err != nil || !r.Success {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, nil)
}
