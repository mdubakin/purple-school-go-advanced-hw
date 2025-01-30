package main

import (
	"orderapi/config"
	"orderapi/internal/cart"
	"orderapi/internal/database"
	"orderapi/internal/product"
	"orderapi/internal/user"
)

func main() {
	cfg := config.MustLoadConfig()

	db, err := database.GetDBConn(&cfg.Postgres)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&user.User{},
		&product.Product{},
		&cart.Cart{},
	)
}
