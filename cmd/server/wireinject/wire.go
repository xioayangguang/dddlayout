//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"layout/apis/http"
	"layout/application/http_handler"
	"layout/application/http_handler/app"
	"layout/application/http_handler/h5"
	"layout/domain/user/model/entities"
	"layout/infrastructure/db/repository"
	_ "layout/pkg/pprof"
)

var HandlerSet = wire.NewSet(
	http_handler.ProviderSet,
	app.ProviderSet,
	app.StructProvider,
	h5.ProviderSet,
	h5.StructProvider,
)

func NewApp() (*gin.Engine, func(), error) {
	panic(wire.Build(
		http.NewServerHTTP,
		repository.ProviderSet,
		entities.ProviderSet,
		HandlerSet,
	))
}
