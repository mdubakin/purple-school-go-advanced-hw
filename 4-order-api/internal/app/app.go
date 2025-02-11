package app

import (
	"log"
	"net/http"
	"orderapi/config"
	"orderapi/internal/controller/database"
	"orderapi/internal/controller/middleware"
	"orderapi/internal/controller/rest"
	repository "orderapi/internal/repository"
	"orderapi/internal/usecase"
)

func Run(cfg *config.Config) {
	db, err := database.GetDBConn(&cfg.Postgres)
	if err != nil {
		panic(err)
	}

	router := http.NewServeMux()

	// repos
	productRepo := repository.NewProductRepository(db)

	// services
	productService := usecase.NewProductService(productRepo)

	// handlers
	rest.NewProductHandler(router, productService)

	chain := middleware.Chain(middleware.WithJSONLogs)

	server := http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: chain(router),
	}

	log.Println("Сервер инициализирован и работает по адресу localhost:" + cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
