package repository

import (
	"go-nunu-server/api"
	"go-nunu-server/internal/model"

	"gorm.io/gorm"
)

type GoodsRepository interface {

	// 查询商品列表
	FindAll() interface{}

	// 根据ID查询商品
	FindByID(id string) interface{}

	CreateGoods(goods *model.Goods) interface{}

	DeleteGoods(id string) (interface{}, error)

	GetDB() *gorm.DB
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
func (r *goodsRepository) CreateGoods(goods *model.Goods) interface{} {
	if err := r.db.Create(&goods).Error; err != nil {
		return err
	}
	return "创建成功" + goods.Name
}

// DeleteGoods 删除商品
func (r *goodsRepository) DeleteGoods(id string) (interface{}, error) {
	var goods model.Goods
	if err := r.db.Where("id = ?", id).First(&goods).Error; err != nil {
		return api.NotFound(id), err
	}
	if err := r.db.Where("id = ?", id).Delete(&goods).Error; err != nil {
		return "删除错误", err
	}
	return "删除成功", nil
}

func (r *goodsRepository) GetDB() *gorm.DB {
	return r.db
}
