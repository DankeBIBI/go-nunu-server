package handler

import (
	"go-nunu-server/api"
	"go-nunu-server/internal/service"

	"github.com/gin-gonic/gin"
)

type UserRouteHandler struct {
	*Handler
	userRouteService service.UserRouteService
}

func NewUserRouteHandler(
	handler *Handler,
	userRouteService service.UserRouteService,
) *UserRouteHandler {
	return &UserRouteHandler{
		Handler:          handler,
		userRouteService: userRouteService,
	}
}

func (h *UserRouteHandler) GetUserRoute(ctx *gin.Context) {
	uid := ctx.PostForm("u_id")
	route, err := h.userRouteService.GetUserRoute(uid)
	if err != nil {
		api.Fail(ctx, err)
	}
	api.Success(ctx, route)

}
