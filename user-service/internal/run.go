package internal

import (
	"fmt"

	"github.com/gene-qxsi/Blog/user-service/config"
	"github.com/gene-qxsi/Blog/user-service/internal/application/service"
	"github.com/gene-qxsi/Blog/user-service/internal/infrastructure/postgres"
	"github.com/gene-qxsi/Blog/user-service/internal/presentation/grpc"
	"github.com/gene-qxsi/Blog/user-service/internal/presentation/http"
)

func Run(config config.Config) error {
	postgresDB, err := postgres.NewPostgresDB(config.Postgres)
	if err != nil {
		return fmt.Errorf("ошибка запуска postgreSQL: %w", err)
	}

	userRepo := postgres.NewUserPostgresRepo(postgresDB)
	userService := service.NewUserService(userRepo)

	userHandlerHTTP := http.NewUserHandler(userService)
	userHandlerGRPC := grpc.NewUserHandler(userService)

	http.RunHTTPServer(config, &http.Handlers{
		UserHandler: userHandlerHTTP,
	})

	grpc.RunGRPCServer(config, &grpc.Handlers{
		UserHandler: userHandlerGRPC,
	})

	return nil
}
