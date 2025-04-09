package app

import (
	"app/internal/config"
	"app/internal/database"
	"app/internal/infrastructure/http"
)

func Run(cfg *config.Config) {
	database.Init(cfg)
	db := database.Get()

	http := http.New(cfg)
	http.Init(db)
	http.Start()
}
