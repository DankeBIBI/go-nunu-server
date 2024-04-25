package repository

import (
	"context"
	"go-nunu-server/internal/model"

	"gorm.io/gorm"
)

type RouteRepository interface {
	GetRoute(ctx context.Context, id int64) (*model.Route, error)

	GetDB() *gorm.DB
}

func NewRouteRepository(
	repository *Repository,
) RouteRepository {
	return &routeRepository{
		Repository: repository,
	}
}

type routeRepository struct {
	*Repository
}

func (r *routeRepository) GetRoute(ctx context.Context, id int64) (*model.Route, error) {
	var route model.Route

	return &route, nil
}

func (r *routeRepository) GetDB() *gorm.DB {
	return r.db
}
