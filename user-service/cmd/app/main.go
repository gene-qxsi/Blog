package main

import (
	"log"

	"github.com/gene-qxsi/Blog/user-service/config"
	app "github.com/gene-qxsi/Blog/user-service/internal"
)

const (
	local = "local"
	prod  = "prod"
)

func main() {
	config, err := config.LoadConfig(local, "yaml", "./config")
	if err != nil {
		log.Fatalln("ошибка загрузки конфигурации:", err)
	}

	if err := app.Run(*config); err != nil {
		log.Fatalln("ошибка запуска приложения:", err)
	}
}
