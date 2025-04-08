package app

import (
	"app/internal/config"
	"app/internal/infrastructure/http"
)

func Run(cfg *config.Config) {
	db := NewDataBase(cfg)

	http := http.New(cfg)
	http.Init(db)
	http.Start()
}
