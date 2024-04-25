package handler

import (
	"fmt"
	"go-nunu-server/api"
	"go-nunu-server/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(
	handler *Handler,
	userService service.UserService,
) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *UserHandler) GetUserList(ctx *gin.Context) {
	phone := ctx.Query("phone")
	user, err := h.userService.GetUserList(phone)
	fmt.Println(err)
	if err != nil {
		api.Fail(ctx, err)
	}
	api.Success(ctx, user)

}

func (h *UserHandler) Login(ctx *gin.Context) {

	var params api.LoginDto
	if err := ctx.ShouldBind(&params); err != nil {
		api.Fail(ctx, err, "数据绑定失败")
		return
	}
	user, err := h.userService.Login(&params)
	if err != nil {
		api.Fail(ctx, user, "登录失败")
		return
	}
	api.Success(ctx, user, "登录成功")
}
