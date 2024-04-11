package handler

import (
	v1 "go-nunu-server/api/v1"
	"go-nunu-server/internal/repository"
	"go-nunu-server/internal/service"

	"github.com/gin-gonic/gin"
)

type GoodsHandler struct {
	*Handler
	goodsService service.GoodsService
	*repository.Repository
}

// NewGoodsHandler 创建商品处理器
func NewGoodsHandler(handler *Handler, goodsService service.GoodsService) *GoodsHandler {
	return &GoodsHandler{
		Handler:      handler,
		goodsService: goodsService,
	}
}

func (h *GoodsHandler) FindAllGoods(ctx *gin.Context) {
	ref := h.goodsService.FindAll(ctx)

	// fmt.Print(goodss)
	v1.HandleSuccess(ctx, ref)
}
