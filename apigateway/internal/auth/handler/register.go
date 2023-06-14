package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/common/cerr"
	"github.com/ikadgzl/inventory/common/proto/auth"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *authHandler) Register(ctx *gin.Context) {
	req := &RegisterRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	r, err := h.authSvc.Register(ctx, &auth.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		cerr.HandleGrpcError(ctx, err)
		return
	}

	// ctx.SetCookie("accessToken", r.Token, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, r.Token)
}
