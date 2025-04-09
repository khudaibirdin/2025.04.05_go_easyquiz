package main

import (
	"app/internal/app"
	"app/internal/config"
)

func main() {
	cfg := config.New("configs/local.yml")
	app.Run(cfg)
}
