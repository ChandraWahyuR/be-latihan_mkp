package config

import (
	"log"

	"github.com/ChandraWahyuR/be-latihan_mkp/common/util"
)

type Config struct {
	PortServer int      `mapstructure:"portServer"`
	Database   Database `mapstructure:"database"`
	JwtSecret  string   `mapstructure:"jwtSecretKey"`
}

type Database struct {
	DbHost string `mapstructure:"host"`
	DbPort int    `mapstructure:"port"`
	DbName string `mapstructure:"name"`
	DbPass string `mapstructure:"password"`
	DbUser string `mapstructure:"username"`
}

func LoadConfig() *Config {
	cfg := &Config{}
	if err := util.BindFromJSON(cfg, "config", "."); err == nil {
		return cfg
	}

	log.Fatal("no config source found")
	return nil
}
