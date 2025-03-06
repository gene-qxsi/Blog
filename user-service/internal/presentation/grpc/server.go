package grpc

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	blogv1 "github.com/gene-qxsi/Blog-api/gen/go"
	"github.com/gene-qxsi/Blog/user-service/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Handlers struct {
	UserHandler *UserHandler
}

func RunGRPCServer(config config.Config, hs *Handlers) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(config.GRPC.Port))
	if err != nil {
		log.Fatal("ошибка создания объекта Listener: %w", err)
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(UnaryLogger))
	blogv1.RegisterUserServer(server, hs.UserHandler)
	reflection.Register(server)

	errCh := make(chan error, 1)
	go func() {
		log.Printf("gRPC сервер запущен на порту %d", config.GRPC.Port)
		errCh <- server.Serve(lis)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		log.Printf("Получен сигнал %s, останавливаем сервер...", sig)
		server.GracefulStop()
	case err := <-errCh:
		log.Fatal("ошибка запуска gRPC сервера: %w", err)
	}
}
