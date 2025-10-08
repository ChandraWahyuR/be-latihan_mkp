package config

import (
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/auth/jwt"
	"github.com/sirupsen/logrus"
)

func NewJWT(log *logrus.Logger, cfg *Config) *jwt.JWTService {
	if cfg.JwtSecret == "" {
		log.Error("Token jwt tidak terbaca atau kosong")
	}
	return jwt.NewJwt(cfg.JwtSecret)
}
