package app

import (
	"database/sql"

	"github.com/ChandraWahyuR/be-latihan_mkp/config"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/auth/jwt"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/delivery/handler"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/delivery/routes"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/repository"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	DB  *sql.DB
	App *gin.Engine
	Log *logrus.Logger
	Cfg *config.Config
	JWT jwt.JWTService
}

func App(config *BootstrapConfig) {
	// repo
	userRepo := repository.NewUserRepository(config.DB, config.Log)
	movieRepo := repository.NewMoviesRepository(config.DB)
	jadwalTayangRepo := repository.NewJadwalTayangRepository(config.DB)

	// usecase
	userUc := usecase.NewUserUsecase(config.Log, userRepo, config.JWT)
	movieUc := usecase.NewMoviesUsecase(movieRepo)
	JadwalTayangUc := usecase.NewJadwalTayangUsecase(jadwalTayangRepo, movieRepo)

	// handler
	userDlv := handler.NewUserHandler(config.JWT, userUc)
	movieDlv := handler.NewMovieHandler(movieUc)
	JadwalTayangDlv := handler.NewJadwalTayangHandler(JadwalTayangUc)

	routeConfig := routes.RouteConfig{
		App:                    config.App,
		UserController:         userDlv,
		MovieController:        movieDlv,
		JadwalTayangController: JadwalTayangDlv,
		JWT:                    config.JWT,
	}

	routeConfig.Setup()
}
