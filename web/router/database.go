package router

import (
	"newsSearchEngine/web/service"

	"github.com/gin-gonic/gin"
)

func InitDatabaseRouter(Router *gin.RouterGroup) {

	databaseRouter := Router.Group("db")
	{
		databaseRouter.GET("list", service.DBS)              // 查看数据库
		databaseRouter.GET("drop", service.DatabaseDrop)     // 删除数据库
		databaseRouter.GET("create", service.DatabaseCreate) // 添加数据库
	}
}
