package main

import (
	"app/internal/adapters/database"
	"app/internal/config"
	"app/internal/entities"
)

func main() {
	cfg := config.New("configs/local.yml")
	database.Init(cfg)
	database.Get().AutoMigrate(
		&entities.User{},
		&entities.Quiz{},
		&entities.Answer{},
		&entities.AnswerVariants{},
		&entities.Question{},
		&entities.Result{},
	)
}
