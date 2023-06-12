package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/common/proto/auth"
)

func (h *authHandler) Validate(ctx *gin.Context) {
	token, err := ctx.Cookie("accessToken")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	r, err := h.authSvc.Validate(ctx, &auth.ValidateRequest{
		Token: token,
	})
	if err != nil || r.Status != http.StatusOK {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r.Error)
		return
	}

	ctx.JSON(http.StatusOK, r.UserId)
}
