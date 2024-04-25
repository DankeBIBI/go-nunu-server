package handler

import (
	"go-nunu-server/api"
	v1 "go-nunu-server/api/v1"
	"go-nunu-server/internal/model"
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
	ref := h.goodsService.FindAll()
	v1.HandleSuccess(ctx, ref)
}
func (h *GoodsHandler) FindByID(ctx *gin.Context) {
	ref := h.goodsService.FindByID(ctx)
	api.Success(ctx, ref)
}
func (h *GoodsHandler) CreateGoods(ctx *gin.Context) {
	var req api.CreateGoodsDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.Fail(ctx, err)
		return
	}
	res := h.goodsService.CreateGoods(&req)
	api.Success(ctx, res, "创建成功"+ctx.PostForm("name"))
}
func (h *GoodsHandler) DeleteGoods(ctx *gin.Context) {
	id := ctx.Query("id")
	res, err := h.goodsService.DeleteGoods(id)
	if err != nil {
		api.Fail(ctx, err, "删除失败："+id)
		return
	}
	api.Success(ctx, res, "删除成功")
}
func (h *GoodsHandler) UpdateGoods(ctx *gin.Context) {
	var req model.Goods
	if err := ctx.ShouldBind(&req); err != nil {
		api.Fail(ctx, err, "数据绑定失败")
		return
	}
	req.Appid = ctx.GetHeader("Appid")
	res, err := h.goodsService.UpdateGoods(&req)
	if err != nil {
		api.Fail(ctx, err, "修改失败")
		return
	}
	api.Success(ctx, res, "修改成功")
}
