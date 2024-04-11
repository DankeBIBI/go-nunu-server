package repository

import (
	"context"
	"fmt"
	"go-nunu-server/internal/model"
)

type GoodsRepository interface {

	// 查询商品列表
	FindAll(ctx context.Context) interface{}
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
func (r *goodsRepository) FindAll(ctx context.Context) interface{} {
	var goods []model.Goods
	list := r.db.Find(&goods)
	fmt.Print(list)
	return &goods
}
