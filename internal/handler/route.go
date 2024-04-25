package handler

import (
	"go-nunu-server/api"
	"go-nunu-server/internal/service"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	*Handler
	routeService service.RouteService
}

func NewRouteHandler(
	handler *Handler,
	routeService service.RouteService,
) *RouteHandler {
	return &RouteHandler{
		Handler:      handler,
		routeService: routeService,
	}
}

func (h *RouteHandler) GetAllRoute(ctx *gin.Context) {
	route, err := h.routeService.GetAllRoute(ctx)
	if err != nil {
		api.Fail(ctx, nil, "查找失败")
	}
	api.Success(ctx, route, "查找成功")
}
