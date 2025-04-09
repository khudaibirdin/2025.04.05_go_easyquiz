package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	_, err := c.AddFunc("*/1 * * * *", func() { fmt.Println("Запуск задачи для подсчета результатов") })
	if err != nil {
		log.Fatal("Ошибка добавления задачи:", err)
	}

	c.Start()
	log.Println("Cron-задачи запущены...")
	select {}
}
