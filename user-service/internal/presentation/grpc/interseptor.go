package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/fatih/color"
	"google.golang.org/grpc"
)

func UnaryLogger(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("[error]: {%s}", err.Error())
	}

	log.Println(color.YellowString(fmt.Sprintf("[server]: {%s}. [method]: {%v}. [result]: {%v}", info.Server, info.FullMethod, resp)))

	return resp, err
}
