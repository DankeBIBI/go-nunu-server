package handler

import (
	"go-nunu-server/api"
	"go-nunu-server/internal/service"

	"github.com/gin-gonic/gin"
)

type UserInfoHandler struct {
	*Handler
	userInfoService service.UserInfoService
}

func NewUserInfoHandler(
	handler *Handler,
	userInfoService service.UserInfoService,
) *UserInfoHandler {
	return &UserInfoHandler{
		Handler:         handler,
		userInfoService: userInfoService,
	}
}

func (h *UserInfoHandler) GetUserInfo(ctx *gin.Context) {

}

func (h *UserInfoHandler) CheckWechatLogin(ctx *gin.Context) {
	var params api.CheckWechatLogin
	if err := ctx.ShouldBind(&params); err != nil {
		api.Fail(ctx, "绑定失败")
	}
	res, err := h.userInfoService.CheckWechatLogin(&params)
	if err != nil {
		api.Fail(ctx, err, "绑定失败")
		return
	}
	api.Success(ctx, res)
}
