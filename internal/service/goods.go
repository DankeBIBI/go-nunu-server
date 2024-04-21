package service

import (
	"fmt"
	"go-nunu-server/api"
	"go-nunu-server/internal/model"
	"go-nunu-server/internal/repository"

	"github.com/gin-gonic/gin"
)

type GoodsService interface {
	//查询商品列表
	FindAll() interface{}
	//根据ID查询商品
	FindByID(ctx *gin.Context) interface{}
	//创建商品
	CreateGoods(goods *api.CreateGoodsDto) interface{}
	//删除商品
	DeleteGoods(id string) (interface{}, error)
	//修改商品信息
	UpdateGoods(goods *model.Goods) (interface{}, error)
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
	newGoods := &model.Goods{
		Name:         goods.Name,
		Price:        float32(goods.Price),
		Source_price: float32(goods.Source_price),
		Stock:        goods.Stock,
		Pic:          goods.Pic,
	}
	res := g.goodsRepo.CreateGoods(newGoods)
	return res
}

// 删除商品
func (g *goodsService) DeleteGoods(id string) (interface{}, error) {
	res, err := g.goodsRepo.DeleteGoods(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

// 修改商品信息
func (g *goodsService) UpdateGoods(goods *model.Goods) (interface{}, error) {
	newGoods := &model.Goods{
		Id:            goods.Id,
		Appid:         goods.Appid,
		Name:          goods.Name,
		Mini_name:     goods.Mini_name,
		Pic_list:      goods.Pic_list,
		Pic:           goods.Pic,
		Price:         goods.Price,
		Quota_num:     goods.Quota_num,
		Virtual_sales: goods.Virtual_sales,
		Sort:          goods.Sort,
		Sales:         goods.Sales,
		Source_price:  goods.Source_price,
	}
	if err := g.goodsRepo.GetDB().Where("id = ?", goods.Id).First(goods).Error; err != nil {
		return "查找失败", err
	}
	if err := g.goodsRepo.GetDB().Save(newGoods).Error; err != nil {
		return "更新失败", err
	}
	fmt.Println(newGoods)
	return "更新成功", nil
}
