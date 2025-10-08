package config

import (
	"database/sql"
	"fmt"
	"log"
)

func InitDatabase(cfg *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Database.DbUser, cfg.Database.DbPass, cfg.Database.DbHost, cfg.Database.DbPort, cfg.Database.DbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfull connect to database on DSN:%s", dsn)
	return db, nil
}
