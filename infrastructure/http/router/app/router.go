package app

import (
	"github.com/gin-gonic/gin"
	middleware2 "layout/infrastructure/middleware"
	"layout/internal/handler/http/app"
	"time"
)

const SignSalt = "T^N5kJDOJ7seK3Z$"

const Timeout = 500 * time.Millisecond

func InitAppRouter(Router *gin.Engine, router *app.Router) {
	ApiRouter := Router.Group("api")
	ApiRouter.Use(middleware2.RequestLog())
	ApiRouter.Use(middleware2.Timeout(Timeout))
	ApiRouter.Use(middleware2.Sign(SignSalt))
	//全局分布式限速
	//ApiRouter.Use(middleware.SpeedLimit())
	//单机限速（如果网关做了ip哈希的优先试用单机限速提高性能）
	ApiRouter.Use(middleware2.TokenLimit())
	ApiRouter.Use(middleware2.Recover())
	//必须登录的路由
	PrivateApiGroup := ApiRouter.Group("")
	PrivateApiGroup.Use(middleware2.MustTokenAuth())
	PrivateApiGroup.Use(middleware2.AccessRecords())
	MustLoginRouter(PrivateApiGroup, router)

	//可以登录也可以不登录的路由
	ShouldLoginApiGroup := ApiRouter.Group("")
	ShouldLoginApiGroup.Use(middleware2.ShouldTokenAuth())
	ShouldLoginApiGroup.Use(middleware2.AccessRecords())
	ShouldLoginRouter(ShouldLoginApiGroup, router)

	//可以不登陆的路由
	PublicApiGroup := ApiRouter.Group("")
	VisitorRouter(PublicApiGroup, router)
}
