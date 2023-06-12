package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/common/proto/user"
)

func (h *userHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	r, err := h.userSvc.GetUser(ctx, &user.GetUserRequest{
		Id: id,
	})
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, r)
}
