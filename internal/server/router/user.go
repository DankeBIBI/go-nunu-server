package router

import (
	"go-nunu-server/internal/handler"

	"github.com/gin-gonic/gin"
)

// 分配路由相关路由
func AssignUserRouter(route *gin.RouterGroup, g *handler.UserHandler) {
	route.GET("/list", g.GetUserList)
	route.POST("/login", g.Login)
}

// 分配微信登录相关路由
func AssignWechatRouter(route *gin.RouterGroup, g *handler.UserInfoHandler) {
	route.POST("/checkLogin", g.CheckWechatLogin)
}
