package app

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(code int, msg string, data interface{}) {
	g.C.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return
}
