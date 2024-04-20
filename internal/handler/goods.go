package handler

import (
	"go-nunu-server/api"
	v1 "go-nunu-server/api/v1"
	"go-nunu-server/internal/repository"
	"go-nunu-server/internal/service"
	"net/http"

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
	ref := h.goodsService.FindAll()
	v1.HandleSuccess(ctx, ref)
}
func (h *GoodsHandler) FindByID(ctx *gin.Context) {
	ref := h.goodsService.FindByID(ctx)
	v1.HandleSuccess(ctx, ref)
}
func (h *GoodsHandler) CreateGoods(ctx *gin.Context) {
	var req api.CreateGoodsDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	res := h.goodsService.CreateGoods(&req)
	v1.HandleSuccess(ctx, res)
}
