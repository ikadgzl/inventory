package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ikadgzl/inventory/apigateway/internal/auth"
	"github.com/ikadgzl/inventory/apigateway/internal/config"
	"github.com/ikadgzl/inventory/apigateway/internal/user"
)

func main() {
	c := config.NewConfig()

	r := gin.Default()

	a := auth.RegisterRoutes(r, c)
	user.RegisterRoutes(r, c, a)

	r.Run(c.Port)
}
