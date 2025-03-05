package domain

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (int, error)
	GetUser(ctx context.Context, id int) (*User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, user User) (*User, error)
}
