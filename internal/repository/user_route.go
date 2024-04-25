package repository

import (
	"context"
	"go-nunu-server/internal/model"

	"gorm.io/gorm"
)

type UserRouteRepository interface {
	GetUserRoute(ctx context.Context, id int64) (*model.UserRoute, error)

	GetDB() *gorm.DB
}

func NewUserRouteRepository(
	repository *Repository,
) UserRouteRepository {
	return &userRouteRepository{
		Repository: repository,
	}
}

type userRouteRepository struct {
	*Repository
}

func (r *userRouteRepository) GetUserRoute(ctx context.Context, id int64) (*model.UserRoute, error) {
	var userRoute model.UserRoute

	return &userRoute, nil
}

func (r *userRouteRepository) GetDB() *gorm.DB {
	return r.db
}
