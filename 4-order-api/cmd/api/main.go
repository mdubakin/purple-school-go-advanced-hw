package main

import (
	"orderapi/config"
	"orderapi/internal/app"
)

func main() {
	cfg := config.MustLoadConfig()
	app.Run(cfg)
}
