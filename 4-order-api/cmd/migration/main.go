package main

import (
	"orderapi/config"
	"orderapi/internal/controller/database"
	"orderapi/internal/model"
)

func main() {
	cfg := config.MustLoadConfig()

	db, err := database.GetDBConn(&cfg.Postgres)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&model.Cart{},
		&model.User{},
		&model.Product{},
	)
}
