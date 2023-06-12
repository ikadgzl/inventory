package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/common/proto/auth"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *authHandler) Register(ctx *gin.Context) {
	req := &RegisterRequest{}
	if err := ctx.Bind(req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	r, err := h.authSvc.Register(ctx, &auth.RegisterRequest{
		Username: req.Email,
		Password: req.Password,
	})
	if err != nil || r.Status != http.StatusOK {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r.Error)
		return
	}

	// TODO: correct values
	ctx.SetCookie("accessToken", r.Token, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, nil)
}
