package app

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler/http/app"
)

func ShouldLoginRouter(Router *gin.RouterGroup, router *app.Router) {
	{
		indexRouter := Router.Group("user")
		_ = indexRouter
	}
}
