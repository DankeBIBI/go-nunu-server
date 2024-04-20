package service

import (
	"go-nunu-server/api"
	"go-nunu-server/internal/repository"

	"github.com/gin-gonic/gin"
)

type GoodsService interface {
	FindAll() interface{}
	FindByID(ctx *gin.Context) interface{}
	CreateGoods(goods *api.CreateGoodsDto) interface{}
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

// 查询商品列表
func (g *goodsService) FindAll() interface{} {
	goods := g.goodsRepo.FindAll()
	return goods
}

// 根据ID查询商品
func (g *goodsService) FindByID(ctx *gin.Context) interface{} {
	id := ctx.PostForm("id")
	// 调用repository层的方法
	goods := g.goodsRepo.FindByID(id)
	return goods
}

// 创建商品
func (g *goodsService) CreateGoods(goods *api.CreateGoodsDto) interface{} {
	res := g.goodsRepo.CreateGoods(goods)
	return res
}
