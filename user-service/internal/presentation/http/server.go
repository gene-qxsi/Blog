package http

import (
	"log"
	"net/http"

	"github.com/gene-qxsi/Blog/user-service/config"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	UserHandler *UserHandler
}

func RunHTTPServer(config config.Config, handlers *Handlers) {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		handlers.UserHandler.RegisterUserRoutes(v1)
	}

	var err error

	go func() {
		if err = http.ListenAndServe(config.HTTP.Port, router); err != nil {
			log.Println("ошибка запуска сервера:", err.Error())
		}
	}()

	if err == nil {
		log.Println("Сервер успешно запушен на порту", config.HTTP.Port)
	}
}
