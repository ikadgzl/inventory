package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/apigateway/internal/auth"
	"github.com/ikadgzl/inventory/apigateway/internal/config"
	"github.com/ikadgzl/inventory/apigateway/internal/user/handler"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, a auth.AuthMiddleware) {
	userHandler := handler.NewUserHandler(NewUserServiceClient(c))

	routes := r.Group("/users")
	routes.Use(a.AuthN)
	routes.GET("/:id", userHandler.GetUser)
	routes.GET("", userHandler.ListUsers)
	routes.PUT("/:id", userHandler.UpdateUser)
	routes.DELETE("/:id", userHandler.DeleteUser)
}
