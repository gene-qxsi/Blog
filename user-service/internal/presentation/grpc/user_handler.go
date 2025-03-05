package grpc

import (
	"context"
	"fmt"

	blogv1 "github.com/gene-qxsi/Blog-api/gen/go"
	"github.com/gene-qxsi/Blog/user-service/internal/application/service"
	"github.com/gene-qxsi/Blog/user-service/internal/presentation/dto"
)

type UserHandler struct {
	srv *service.UserService
	blogv1.UnimplementedUserServer
}

func NewUserHandler(srv *service.UserService) *UserHandler {
	return &UserHandler{srv: srv}
}

func (h *UserHandler) GetUser(ctx context.Context, req *blogv1.GetUserRequest) (*blogv1.GetUserResponse, error) {
	const op = "user-service>internal>presentation>grpc>user_handler.go>GetUser()"

	user, err := h.srv.GetUser(ctx, int(req.Id))
	if err != nil {
		return nil, fmt.Errorf("ошибка: %s. operation: %s", err.Error(), op)
	}

	return &blogv1.GetUserResponse{
		Id:       int64(user.ID()),
		Email:    user.Email(),
		Password: user.Password(),
	}, nil
}

func (h *UserHandler) GetUsers(ctx context.Context, req *blogv1.Empty) (*blogv1.UserList, error) {
	const op = "user-service>internal>presentation>grpc>user_handler.go>GetUser()"

	usersReq, err := h.srv.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка: %s. operation: %s", err.Error(), op)
	}

	var users []*blogv1.UserEntity
	for _, user := range usersReq {
		user := blogv1.UserEntity{
			Id:       int32(user.ID()),
			Email:    user.Email(),
			Password: user.Password(),
		}
		users = append(users, &user)
	}
	return &blogv1.UserList{Users: users}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *blogv1.UpdateUserRequest) (*blogv1.UpdateUserResponse, error) {
	const op = "user-service>internal>presentation>grpc>user_handler.go>UpdateUser()"

	user, err := h.srv.UpdateUser(ctx, int(req.Id), dto.UserRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, fmt.Errorf("ошибка: %s. operation: %s", err.Error(), op)
	}

	return &blogv1.UpdateUserResponse{
		Id: int64(user.ID()),
		// Username: user.Username(),
		Email:    user.Email(),
		Password: user.Password(),
	}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *blogv1.CreateUserRequest) (*blogv1.CreateUserResponse, error) {
	const op = "user-service>internal>presentation>grpc>user_handler.go>CreateUser()"

	id, err := h.srv.CreateUser(ctx, dto.UserRequest{
		// Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка: %s. operation: %s", err.Error(), op)
	}

	return &blogv1.CreateUserResponse{
		Id: int64(id),
	}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *blogv1.DeleteUserRequest) (*blogv1.Empty, error) {
	const op = "user-service>internal>presentation>grpc>user_handler.go>DeleteUser()"

	err := h.srv.DeleteUser(ctx, int(req.Id))
	if err != nil {
		return nil, fmt.Errorf("ошибка: %s. operation: %s", err.Error(), op)
	}

	return &blogv1.Empty{}, nil
}
