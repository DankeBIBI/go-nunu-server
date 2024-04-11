package web

// package handler

// import (
// 	v1 "go-nunu-server/api/v1"
// 	"go-nunu-server/internal/service"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type GoodsHandler struct {
// 	*Handler
// 	goodsService service.GoodsService
// }

// // NewGoodsHandler 创建商品处理器
// func NewGoodsHandler(handler *Handler, goodsService service.GoodsService) *GoodsHandler {
// 	return &GoodsHandler{
// 		Handler:      handler,
// 		goodsService: goodsService,
// 	}
// }

// func (h *GoodsHandler) FindAllGoods(ctx *gin.Context) {
// 	if err := h.goodsService.FindAll(ctx); err != nil {
// 		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
// 		return
// 	}
// 	v1.HandleSuccess(ctx, nil)
// }
