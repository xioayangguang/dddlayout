package h5

import (
	"github.com/gin-gonic/gin"
	"layout/internal/handler"
)

func ShouldLoginRouter(Router *gin.RouterGroup, userHandler handler.UserHandler) {
	{
		indexRouter := Router.Group("match")
		_ = indexRouter
	}
}
