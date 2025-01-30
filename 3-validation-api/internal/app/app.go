package app

import (
	"log"
	"net/http"
	"validation/config"
	"validation/internal/controller"
	"validation/internal/usecase/hash"
	"validation/internal/usecase/repo"
)

func Run(cfg *config.Config) error {
	mux := http.NewServeMux()

	repo, err := repo.NewLocalJSONRepo(cfg.Database.LocalJSONConfig.Path)
	if err != nil {
		return err
	}
	emailHashService := hash.NewHashService(cfg, repo)
	controller.NewVerifyHandler(mux, controller.VerifyHandlerDeps{HashService: *emailHashService})

	server := http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: mux,
	}

	log.Printf("Сервер работает на порту %v\n", cfg.Server.Port)
	return server.ListenAndServe()
}
