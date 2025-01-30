package main

import (
	"log"
	"validation/config"
	"validation/internal/app"
)

func main() {
	cfg := config.MustLoadConfig()

	if err := app.Run(cfg); err != nil {
		log.Panicln(err)
	}
}
