package repository

import (
	"go-nunu-server/api"
	"go-nunu-server/internal/model"
)

type GoodsRepository interface {

	// 查询商品列表
	FindAll() interface{}

	// 根据ID查询商品
	FindByID(id string) interface{}

	CreateGoods(goods *api.CreateGoodsDto) string
}
type goodsRepository struct {
	*Repository
}

// NewGoodsRepository 创建一个新的GoodsRepository实例
func NewGoodsRepository(r *Repository) GoodsRepository {
	return &goodsRepository{
		Repository: r,
	}
}

// FindAll 查找所有商品
func (r *goodsRepository) FindAll() interface{} {
	var goods []model.Goods
	if err := r.db.Where("appid = ?", "10").Find(&goods).Error; err != nil {
		return "查询错误"
	}
	return &goods
}

// FindByID 根据ID查找商品
func (r *goodsRepository) FindByID(id string) interface{} {
	var goods model.Goods
	if err := r.db.Where("id = ?", id).First(&goods).Error; err != nil {
		return "查询错误"
	}
	return &goods
}

// CreateGoods 创建商品
func (r *goodsRepository) CreateGoods(goods *api.CreateGoodsDto) string {
	if err := r.db.Create(&goods).Error; err != nil {
		return "创建失败" + goods.Name
	}
	return "创建成功" + goods.Name
}
