package postgres

import (
	"context"
	"fmt"

	"github.com/gene-qxsi/Blog-user/internal/domain"
	"gorm.io/gorm"
)

type userDB struct {
	ID       int    `gorm:"type:int;primaryKey"`
	Email    string `gorm:"type:varchar(64);unique index"`
	Password string `gorm:"type:text"`
}

func (userDB) TableName() string {
	return "users"
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserPostgresRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r UserRepo) CreateUser(ctx context.Context, user domain.User) (int, error) {
	// const op = "user-service>internal>infrastructure>http>user_create.go>CreateUser()"

	userDB := userDB{
		Email:    user.Email(),
		Password: user.Password(),
	}

	if err := r.db.WithContext(ctx).Debug().Create(&userDB).Error; err != nil {
		return 0, err
	}

	return userDB.ID, nil
}

func (r UserRepo) GetUser(ctx context.Context, id int) (*domain.User, error) {
	// const op = "user-service>internal>infrastructure>http>user_create.go>CreateUser()"

	var userDB userDB
	if err := r.db.WithContext(ctx).Debug().First(&userDB, id).Error; err != nil {
		return nil, err
	}

	user, err := domain.NewUser(userDB.ID, userDB.Email, userDB.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepo) DeleteUser(ctx context.Context, id int) error {
	// const op = "user-service>internal>infrastructure>http>user_create.go>DeleteUser()"

	result := r.db.WithContext(ctx).Debug().Delete(&userDB{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("user not exists")
	}

	return nil
}

func (r UserRepo) UpdateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	// const op = "user-service>internal>infrastructure>http>user_create.go>UpdateUser()"
	userDB := userDB{
		ID:       user.ID(),
		Email:    user.Email(),
		Password: user.Password(),
	}

	if err := r.db.WithContext(ctx).Debug().Model(&userDB).Updates(userDB).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).First(&userDB, userDB.ID).Error; err != nil {
		return nil, err
	}

	updatedUser, err := domain.NewUser(userDB.ID, userDB.Email, userDB.Password)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
