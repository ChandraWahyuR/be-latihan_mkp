package routes

import (
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/auth/jwt"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/delivery/handler"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App                    *gin.Engine
	UserController         *handler.UserHandler
	MovieController        *handler.MovieHandler
	JadwalTayangController *handler.JadwalTayangHandler
	JWT                    jwt.JWTService
}

func (c *RouteConfig) Setup() {
	c.App.Use(gin.Logger())

	v1 := c.App.Group("/api/v1")

	c.setupPublicRoutes(v1)
	c.setupAuthRoutes(v1)
}

func (c *RouteConfig) setupPublicRoutes(router *gin.RouterGroup) {
	// Login & Register
	router.POST("/register", c.UserController.Register)
	router.POST("/login", c.UserController.Login)
}

func (c *RouteConfig) setupAuthRoutes(router *gin.RouterGroup) {
	auth := router.Group("").Use(middleware.NewAuth(c.JWT))

	{
		auth.GET("/movie", c.MovieController.GetAllMovies)
		auth.GET("/movie/:id", c.MovieController.GetMovieDetailByID)
		auth.POST("/movie", c.MovieController.CreateMovie)
		auth.PUT("/movie/:id", c.MovieController.UpdateMovie)
		auth.DELETE("/movie/:id", c.MovieController.DeleteMovie)
	}

	{
		auth.GET("/jadwal-tayang", c.JadwalTayangController.GetAllJadwalTayang)
		auth.GET("/jadwal-tayang/:id", c.JadwalTayangController.GetJadwalTayangByID)
		auth.POST("/jadwal-tayang", c.JadwalTayangController.CreateJadwalTayang)
		auth.PUT("/jadwal-tayang/:id", c.JadwalTayangController.UpdateJadwalTayang)
		auth.DELETE("/jadwal-tayang/:id", c.JadwalTayangController.DeleteJadwalTayang)
	}

}
