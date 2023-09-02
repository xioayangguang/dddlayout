package h5

import (
	"github.com/gin-gonic/gin"
	middleware2 "layout/infrastructure/middleware"
	"layout/internal/handler/http/h5"
	"time"
)

const SignSalt = "bWAOoXvIqxeiqk6*"

const Timeout = 500 * time.Millisecond

func InitH5Router(Router *gin.Engine, router *h5.Router) {
	H5Router := Router.Group("h5")
	H5Router.Use(middleware2.RequestLog())
	H5Router.Use(middleware2.Timeout(Timeout))
	H5Router.Use(middleware2.CORSMiddleware())
	H5Router.Use(middleware2.Sign(SignSalt))
	H5Router.Use(middleware2.SpeedLimit())
	H5Router.Use(middleware2.Recover())
	//必须登录的路由
	PrivateApiGroup := H5Router.Group("")
	PrivateApiGroup.Use(middleware2.MustTokenAuth())
	PrivateApiGroup.Use(middleware2.AccessRecords())
	MustLoginRouter(PrivateApiGroup, router)
	//可以登录也可以不登录的路由
	ShouldLoginApiGroup := H5Router.Group("")
	ShouldLoginApiGroup.Use(middleware2.ShouldTokenAuth())
	ShouldLoginApiGroup.Use(middleware2.AccessRecords())
	ShouldLoginRouter(ShouldLoginApiGroup, router)
	//可以不登陆的路由
	PublicApiGroup := H5Router.Group("")
	VisitorRouter(PublicApiGroup, router)
}
