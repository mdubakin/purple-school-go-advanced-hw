package app

import (
	"orderapi/config"
	"orderapi/internal/database"
)

func Run(cfg *config.Config) {
	_, err := database.GetDBConn(&cfg.Postgres)
	if err != nil {
		panic(err)
	}
}
