package app

import (
	"github.com/gin-gonic/gin"
	"layout/application/http_handler/app"
)

func ShouldLoginRouter(Router *gin.RouterGroup, router *app.Router) {
	{
		indexRouter := Router.Group("user")
		_ = indexRouter
	}
}
