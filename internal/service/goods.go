package service

import (
	"context"
	"go-nunu-server/internal/repository"
)

type GoodsService interface {
	FindAll(ctx context.Context) interface{}
}
type goodsService struct {
	goodsRepo repository.GoodsRepository
	*Service
}

func NewGoodsService(service *Service, goodsRepo repository.GoodsRepository) GoodsService {
	return &goodsService{
		goodsRepo: goodsRepo,
		Service:   service,
	}
}

func (g *goodsService) FindAll(ctx context.Context) interface{} {
	goods := g.goodsRepo.FindAll(ctx)
	return goods
}
