package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/common/proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *authHandler) Login(ctx *gin.Context) {
	req := &LoginRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	r, err := h.authSvc.Login(ctx, &auth.LoginRequest{
		Username: req.Email,
		Password: req.Password,
	})
	if err != nil {
		grpcErr, ok := status.FromError(err)
		fmt.Printf("error code: %v+\n", grpcErr)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		switch grpcErr.Code() {
		case codes.InvalidArgument:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": grpcErr.Message(),
			})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": grpcErr.Message(),
			})
			return
		}
	}

	// TODO: correct values
	ctx.SetCookie("accessToken", r.Token, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, nil)
}
