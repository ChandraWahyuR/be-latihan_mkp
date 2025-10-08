package main

import (
	"fmt"
	"net/http"

	"github.com/ChandraWahyuR/be-latihan_mkp/app"
	"github.com/ChandraWahyuR/be-latihan_mkp/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	cfg := config.LoadConfig()
	logger := logrus.New()
	jwt := config.NewJWT(logger, cfg)
	db, err := config.InitDatabase(cfg)
	if err != nil {
		logger.Fatal("failed connect to database:", err)
		return
	}

	bootstrap := &app.BootstrapConfig{
		App: router,
		DB:  db,
		Log: logger,
		JWT: *jwt,
		Cfg: cfg,
	}
	app.App(bootstrap)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	port := fmt.Sprintf(":%d", cfg.PortServer)
	fmt.Println("listen to port", port)
	err = router.Run(port)
	if err != nil {
		fmt.Println("Server tidak bisa dijalankan:", err)
	}
}
