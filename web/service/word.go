package service

import (
	"newsSearchEngine/engine"

	"github.com/gin-gonic/gin"
)

func WordCut(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		ResponseErrorWithMsg(c, "请输入关键字")
		return
	}
	r := engine.Tokenizer.Cut(q)
	ResponseSuccessWithData(c, r)
}
