package service

import (
	"context"
	"go-nunu-server/internal/model"
	"go-nunu-server/internal/repository"
)

type RouteService interface {
	GetAllRoute(ctx context.Context) (interface{}, error)
}

func NewRouteService(service *Service, routeRepository repository.RouteRepository) RouteService {
	return &routeService{
		Service:         service,
		routeRepository: routeRepository,
	}
}

type routeService struct {
	*Service
	routeRepository repository.RouteRepository
}

func (s *routeService) GetAllRoute(ctx context.Context) (interface{}, error) {
	var route []model.Route
	if err := s.routeRepository.GetDB().Find(&route).Error; err != nil {
		return nil, err
	}
	return &route, nil
}
