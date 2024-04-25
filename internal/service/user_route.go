package service

import (
	"go-nunu-server/internal/model"
	"go-nunu-server/internal/repository"
)

type UserRouteService interface {
	GetUserRoute(id string) (interface{}, error)
}

func NewUserRouteService(service *Service, userRouteRepository repository.UserRouteRepository) UserRouteService {
	return &userRouteService{
		Service:             service,
		userRouteRepository: userRouteRepository,
	}
}

type userRouteService struct {
	*Service
	userRouteRepository repository.UserRouteRepository
}

// 用户路由
func (s *userRouteService) GetUserRoute(id string) (interface{}, error) {
	var userRoute []model.UserRoute
	if err := s.userRouteRepository.GetDB().Preload("Route").Find(&userRoute, "u_id = ?", id).Error; err != nil {
		return "查找错误", err
	}
	return userRoute, nil
}
