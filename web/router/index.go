package router

import (
	"newsSearchEngine/web/service"

	"github.com/gin-gonic/gin"
)

func InitIndexRouter(Router *gin.RouterGroup) {

	indexRouter := Router.Group("index")
	{
		indexRouter.POST("", service.AddIndex)           // 添加单条索引
		indexRouter.POST("batch", service.BatchAddIndex) // 批量添加索引
		indexRouter.POST("remove", service.RemoveIndex)  // 删除索引

		indexRouter.POST("query", service.Query) //query search
		indexRouter.GET("gc", service.GC)
	}
}
