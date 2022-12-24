package service

import (
	"newsSearchEngine/engine"
	"newsSearchEngine/engine/model"
	"runtime"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	var request = &model.SearchRequest{
		Database: c.Query("database"),
	}
	if request.Database == "" {
		request.Database = "default"
	}
	if err := c.ShouldBind(&request); err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	//调用搜索
	r, err := engine.OpEngine.MultiSearch(request)
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
	} else {
		ResponseSuccessWithData(c, r)
	}
}

// AddIndex 添加索引
func AddIndex(c *gin.Context) {
	document := &model.IndexDoc{}
	if err := c.ShouldBindJSON(&document); err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}
	dbName := c.Query("database")
	if dbName == "" {
		ResponseErrorWithMsg(c, "database is empty")
		return
	}
	err := engine.OpEngine.IndexDocument(document)
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}

	ResponseSuccessWithData(c, nil)
}

// BatchAddIndex 批量添加索引
func BatchAddIndex(c *gin.Context) {
	documents := make([]*model.IndexDoc, 0)
	if err := c.BindJSON(&documents); err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}

	dbName := c.Query("database")
	if dbName == "" {
		ResponseErrorWithMsg(c, "database is empty")
		return
	}

	for _, doc := range documents {
		if err := engine.OpEngine.IndexDocument(doc); err != nil {
			ResponseErrorWithMsg(c, err.Error())
			return
		}
	}

	ResponseSuccess(c)
}

// RemoveIndex 删除索引
func RemoveIndex(c *gin.Context) {
	removeIndexModel := &model.RemoveIndexModel{}
	if err := c.BindJSON(&removeIndexModel); err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}

	dbName := c.Query("database")
	if dbName == "" {
		ResponseErrorWithMsg(c, "database is empty")
		return
	}

	if err := engine.OpEngine.RemoveIndex(removeIndexModel.Id); err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}

	ResponseSuccess(c)
}

func GC(c *gin.Context) {
	runtime.GC()
	ResponseSuccess(c)
}
