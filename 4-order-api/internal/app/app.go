package app

import (
	"log"
	"net/http"
	"orderapi/config"
	"orderapi/internal/database"
	"orderapi/internal/product"
)

func Run(cfg *config.Config) {
	db, err := database.GetDBConn(&cfg.Postgres)
	if err != nil {
		panic(err)
	}

	router := http.NewServeMux()

	// repos
	productRepo := product.NewProductRepository(db)

	// services
	productService := product.NewProductService(productRepo)

	// handlers
	product.NewProductHandler(router, productService)

	log.Println("Сервер инициализирован и работает по адресу localhost:" + cfg.Server.Port)
	if err := http.ListenAndServe("localhost:"+cfg.Server.Port, router); err != nil {
		panic(err)
	}
}
