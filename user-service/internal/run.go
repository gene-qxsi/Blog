package internal

import (
	"fmt"

	"github.com/gene-qxsi/Blog-user/config"
	"github.com/gene-qxsi/Blog-user/internal/application/service"
	"github.com/gene-qxsi/Blog-user/internal/infrastructure/postgres"
	httpH "github.com/gene-qxsi/Blog-user/internal/presentation/http"
	"github.com/gin-gonic/gin"
)

func Run(config config.Config) (*gin.Engine, error) {
	postgresDB, err := postgres.NewPostgresDB(config.Postgres)
	if err != nil {
		return nil, fmt.Errorf("ошибка запуска postgreSQL: %w", err)
	}

	userRepo := postgres.NewUserPostgresRepo(postgresDB)
	userService := service.NewUserService(userRepo)

	userHandler := httpH.NewUserHandler(userService)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		userHandler.RegisterUserRoutes(v1)
	}

	return router, nil
}
