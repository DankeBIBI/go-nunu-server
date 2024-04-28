package repository

import (
	"gorm.io/gorm"
)

type AppRepository interface {
	GetDB() *gorm.DB
}

func NewAppRepository(
	repository *Repository,
) AppRepository {
	return &appRepository{
		Repository: repository,
	}
}

type appRepository struct {
	*Repository
}

func (r *appRepository) GetDB() *gorm.DB {
	return r.db
}
