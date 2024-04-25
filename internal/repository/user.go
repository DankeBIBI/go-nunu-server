package repository

import (
	"context"
	"go-nunu-server/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(ctx context.Context, id int64) (*model.User, error)

	GetDB() *gorm.DB
}

func NewUserRepository(
	repository *Repository,
) UserRepository {
	return &userRepository{
		Repository: repository,
	}
}

type userRepository struct {
	*Repository
}

func (r *userRepository) GetUser(ctx context.Context, id int64) (*model.User, error) {
	var user model.User

	return &user, nil
}

func (r *userRepository) GetDB() *gorm.DB {
	return r.db
}
