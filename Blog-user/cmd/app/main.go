package main

import (
	"log"
	"net/http"

	"github.com/gene-qxsi/Blog-user/config"
	app "github.com/gene-qxsi/Blog-user/internal"
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

	router, err := app.Run(*config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Сервер успешно запушен на порту", config.HTTP.Port)
	if err := http.ListenAndServe(config.HTTP.Port, router); err != nil {
		log.Println("ошибка запуска сервера:", err.Error())
	}

}
