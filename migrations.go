package main

import (
	"app/internal/config"
	"app/internal/database"
	"app/internal/entities"
)

func main() {
	cfg := config.New("configs/local.yml")
	database.Init(cfg)
	database.Get().AutoMigrate(
		&entities.User{},
	)
}

