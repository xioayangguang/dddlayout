//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"layout/internal/handler/http"
	"layout/internal/handler/http/app"
	"layout/internal/handler/http/h5"
	"layout/internal/repository"
	"layout/internal/router"
	"layout/internal/service"
	_ "layout/pkg/pprof"
)

var HandlerSet = wire.NewSet(
	http.ProviderSet,
	app.ProviderSet,
	app.StructProvider,
	h5.ProviderSet,
	h5.StructProvider,
)

func NewApp() (*gin.Engine, func(), error) {
	panic(wire.Build(
		router.NewServerHTTP,
		repository.ProviderSet,
		service.ProviderSet,
		HandlerSet,
	))
}
