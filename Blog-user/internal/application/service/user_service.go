package service

import (
	"context"

	"github.com/gene-qxsi/Blog-user/internal/domain"
	"github.com/gene-qxsi/Blog-user/internal/presentation/dto"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) *UserService {
	return &UserService{repo: userRepo}
}

func (s UserService) CreateUser(ctx context.Context, req dto.UserRequest) (int, error) {
	// const op = "user-service>internal>application>service>user_service.go>CreateUser()"
	user, err := domain.CreateUser(req.Email, req.Password)
	if err != nil {
		return 0, err
	}

	id, err := s.repo.CreateUser(ctx, *user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s UserService) GetUser(ctx context.Context, id int) (*domain.User, error) {
	// const op = "user-service>internal>application>service>user_service.go>GetUserByID()"
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) DeleteUser(ctx context.Context, id int) error {
	// const op = "user-service>internal>application>service>user_service.go>DeleteUser()"
	return s.repo.DeleteUser(ctx, id)
}

func (s UserService) UpdateUser(ctx context.Context, userID int, req dto.UserRequest) (*domain.User, error) {
	// const op = "user-service>internal>application>service>user_service.go>UpdateUser()"
	user, err := domain.NewUser(userID, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return s.repo.UpdateUser(ctx, *user)
}
