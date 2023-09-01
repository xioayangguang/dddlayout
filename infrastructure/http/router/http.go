package router

import (
	"github.com/gin-gonic/gin"
	"layout/infrastructure/global"
	"layout/infrastructure/router/app"
	"layout/infrastructure/router/h5"
	apphandler "layout/internal/handler/http/app"
	h5handler "layout/internal/handler/http/h5"
	"layout/pkg/rotatelogs"
)

func NewServerHTTP(apphandler *apphandler.Router, h5handler *h5handler.Router) *gin.Engine {
	var r *gin.Engine
	if !global.Config.Debug {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		//r.Use(gin.LoggerWithWriter(rotatelogs.GetRotateLogs("output")))
		r.Use(gin.RecoveryWithWriter(rotatelogs.GetRotateLogs("recovery")))
	} else {
		r = gin.Default()
	}
	app.InitAppRouter(r, apphandler)
	h5.InitH5Router(r, h5handler)
	InitApiRouter(r, apphandler, h5handler)
	InitExtraRouter(r)
	return r
}
