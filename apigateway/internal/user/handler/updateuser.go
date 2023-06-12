package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/common/proto/user"
)

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	req := &UpdateUserRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	r, err := h.userSvc.UpdateUser(ctx, &user.UpdateUserRequest{
		Id:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil || !r.Success {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, nil)
}
