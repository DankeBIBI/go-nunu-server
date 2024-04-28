package repository

import (
	"context"
	"go-nunu-server/internal/model"

	"gorm.io/gorm"
)

type UserInfoRepository interface {
	GetUserInfo(ctx context.Context, id int64) (*model.UserInfo, error)
	GetDB() *gorm.DB
}

func NewUserInfoRepository(
	repository *Repository,
) UserInfoRepository {
	return &userInfoRepository{
		Repository: repository,
	}
}

type userInfoRepository struct {
	*Repository
}

func (r *userInfoRepository) GetUserInfo(ctx context.Context, id int64) (*model.UserInfo, error) {
	var userInfo model.UserInfo

	return &userInfo, nil
}

func (r *userInfoRepository) GetDB() *gorm.DB {
	return r.db
}
