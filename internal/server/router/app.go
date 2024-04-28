package router

import (
	"go-nunu-server/internal/handler"

	"github.com/gin-gonic/gin"
)

// 分配项目相关路由
func AssignAppRouter(route *gin.RouterGroup, g *handler.AppHandler) {
	route.GET("/config", g.GetProject)
}
