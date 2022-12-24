package router

import (
	"newsSearchEngine/web/service"

	"github.com/gin-gonic/gin"
)

func InitWordRouter(Router *gin.RouterGroup) {

	wordRouter := Router.Group("word")
	{
		wordRouter.GET("cut", service.WordCut)
	}
}
