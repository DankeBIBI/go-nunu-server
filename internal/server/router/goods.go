package router

import (
	"go-nunu-server/internal/handler"

	"github.com/gin-gonic/gin"
)

// 分配商品相关路由
func AssignGoodsRouter(route *gin.RouterGroup, g *handler.GoodsHandler) {
	route.GET("/getGoodsList", g.FindAllGoods)
	route.POST("/goodsDetails", g.FindByID)
	route.POST("/createGoods", g.CreateGoods)
	route.GET("/deleteGoods", g.DeleteGoods)
	route.POST("/updateGoods", g.UpdateGoods)
}
