package main

import (
	"app/internal/config"
	"app/internal/app"
)

func main() {
	cfg := config.New()

	app.Run(*cfg)
}
