package http

import (
	"github.com/gin-gonic/gin"
	"layout/apis/http/app"
	"layout/apis/http/h5"
	apphandler "layout/application/http_handler/app"
	h5handler "layout/application/http_handler/h5"
	"layout/infrastructure/config"
	"layout/pkg/rotatelogs"
)

func NewServerHTTP(apphandler *apphandler.Router, h5handler *h5handler.Router) *gin.Engine {
	var r *gin.Engine
	if !config.Config.Debug {
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
