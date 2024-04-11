package router

import (
	"go-nunu-server/internal/handler"

	"github.com/gin-gonic/gin"
)

// 分配商品相关路由
func AssignGoodsRouter(route *gin.RouterGroup, g *handler.GoodsHandler) {
	route.GET("/goods", g.FindAllGoods)
}
