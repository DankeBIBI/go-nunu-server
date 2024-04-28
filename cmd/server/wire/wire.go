//go:build wireinject
// +build wireinject

package wire

import (
	"go-nunu-server/internal/handler"
	"go-nunu-server/internal/repository"
	"go-nunu-server/internal/server"
	"go-nunu-server/internal/service"
	"go-nunu-server/pkg/app"
	"go-nunu-server/pkg/jwt"
	"go-nunu-server/pkg/log"
	"go-nunu-server/pkg/server/http"
	"go-nunu-server/pkg/sid"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewGoodsRepository,
	repository.NewRouteRepository,
	repository.NewUserRouteRepository,
	repository.NewUserInfoRepository,
	repository.NewAppRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewGoodsService,
	service.NewRouteService,
	service.NewUserRouteService,
	service.NewUserInfoService,
	service.NewAppService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewGoodsHandler,
	handler.NewRouteHandler,
	handler.NewUserRouteHandler,
	handler.NewUserInfoHandler,
	handler.NewAppHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
