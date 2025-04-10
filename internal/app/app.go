package app

import (
	"app/internal/adapters/database"
	"app/internal/config"
	"app/internal/infrastructure/http"
)

func Run(cfg *config.Config) {
	database.Init(cfg)
	db := database.Get()

	http := http.New(cfg)
	http.Init(db)
	http.Start()
}
