package app

import (
	"github.com/gin-gonic/gin"
	"layout/application/http_handler/app"
)

func VisitorRouter(Router *gin.RouterGroup, router *app.Router) {
	{
		indexRouter := Router.Group("user")
		indexRouter.POST("/login", router.AppUser.Login)
	}
}
