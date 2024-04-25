package router

import (
	"go-nunu-server/internal/handler"

	"github.com/gin-gonic/gin"
)

// 分配路由相关路由
func AssignRouteRouter(route *gin.RouterGroup, g *handler.RouteHandler, u *handler.UserRouteHandler) {
	route.GET("/GetAllRoute", g.GetAllRoute)
	route.POST("/GetUserRoute", u.GetUserRoute)
}
