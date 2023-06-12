package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/apigateway/internal/auth/handler"
	"github.com/ikadgzl/inventory/apigateway/internal/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) AuthMiddleware {
	authSvc := AuthServiceClient(c)
	authHandler := handler.NewAuthHandler(authSvc)

	routes := r.Group("/auth")
	routes.POST("/login", authHandler.Login)
	routes.POST("/register", authHandler.Register)

	authMiddleware := NewAuthMiddleware(authSvc)
	return authMiddleware
}
