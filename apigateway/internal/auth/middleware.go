package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/common/cerr"
	"github.com/ikadgzl/inventory/common/proto/auth"
)

type AuthMiddleware interface {
	AuthN(ctx *gin.Context)
}

type authMiddleware struct {
	authSvc auth.AuthServiceClient
}

func NewAuthMiddleware(authSvc auth.AuthServiceClient) AuthMiddleware {
	return &authMiddleware{
		authSvc: authSvc,
	}
}

func (m *authMiddleware) AuthN(ctx *gin.Context) {
	token, err := ctx.Cookie("accessToken")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	r, err := m.authSvc.Validate(ctx, &auth.ValidateRequest{
		Token: token,
	})
	if err != nil {
		cerr.HandleGrpcError(ctx, err)
		return
	}

	ctx.Set("userId", r.UserId)
	ctx.Next()
}
