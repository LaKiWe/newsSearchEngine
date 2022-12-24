package service

import (
	"newsSearchEngine/engine"

	"github.com/gin-gonic/gin"
)

func DatabaseDrop(c *gin.Context) {
	dbName := c.Query("database")
	if dbName == "" {
		ResponseErrorWithMsg(c, "database is empty")
		return
	}

	if err := engine.OpEngine.Drop(); err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}

	ResponseSuccessWithData(c, "删除成功")
}

func DatabaseCreate(c *gin.Context) {
	dbName := c.Query("database")
	if dbName == "" {
		ResponseErrorWithMsg(c, "database is empty")
		return
	}

	engine.GetDataBase(dbName)
	ResponseSuccessWithData(c, "创建成功")
}

// database search
func DBS(c *gin.Context) {
	ResponseSuccessWithData(c, map[string]*engine.Engine{engine.OpEngine.DatabaseName: engine.OpEngine})
}
