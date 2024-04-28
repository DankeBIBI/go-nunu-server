package handler

import (
	"go-nunu-server/api"
	"go-nunu-server/internal/service"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	*Handler
	appService service.AppService
}

func NewAppHandler(
	handler *Handler,
	appService service.AppService,
) *AppHandler {
	return &AppHandler{
		Handler:    handler,
		appService: appService,
	}
}

// 获取项目列表
func (h *AppHandler) GetProject(ctx *gin.Context) {
	appid := ctx.Query("appid")
	// api.Fail(ctx, appid)
	// return
	res, err := h.appService.GetProject(appid)
	if err != nil {
		api.Fail(ctx, err)
		return
	}
	api.Success(ctx, res)
}
