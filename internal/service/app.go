package service

import (
	"go-nunu-server/internal/model"
	"go-nunu-server/internal/repository"
)

type AppService interface {
	GetProject(appid string) (interface{}, error)
}

func NewAppService(service *Service, appRepository repository.AppRepository) AppService {
	return &appService{
		Service:       service,
		appRepository: appRepository,
	}
}

type appService struct {
	*Service
	appRepository repository.AppRepository
}

// GetProject 获取项目信息
func (s *appService) GetProject(appid string) (interface{}, error) {
	var app model.App
	if err := s.appRepository.GetDB().Where("id = ?", appid).Find(&app).Error; err != nil {
		return nil, err
	}
	var theme []model.ProjectTheme
	if err := s.appRepository.GetDB().Where("appid = ?", appid).Find(&theme).Error; err != nil {
		return nil, err
	}
	var setting model.ProjectSetting
	if err := s.appRepository.GetDB().Where("appid = ?", appid).Find(&setting).Error; err != nil {
		return nil, err
	}
	app.Theme = theme
	app.Project_name = setting.Project_name
	app.Index_ad = setting.Index_ad
	app.Index_tip = setting.Index_tip
	return app, nil
}
