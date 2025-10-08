package routes

import (
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/auth/jwt"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App            *gin.Engine
	UserController *handler.UserHandler
	JWT            jwt.JWTService
}

func (c *RouteConfig) Setup() {
	c.App.Use(gin.Logger())

	v1 := c.App.Group("/api/v1")

	c.setupPublicRoutes(v1)
}

func (c *RouteConfig) setupPublicRoutes(router *gin.RouterGroup) {
	// Login & Register
	router.POST("/register", c.UserController.Register)
	router.POST("/login", c.UserController.Login)
}
