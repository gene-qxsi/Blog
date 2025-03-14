package domain

import (
	"context"

	"github.com/gene-qxsi/Blog/user-service/internal/presentation/dto"
)

type UserService interface {
	CreateUser(ctx context.Context, req dto.UserRequest) (int, error)
	GetUser(ctx context.Context, id int) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, userID int, req dto.UserRequest) (*User, error)
}
