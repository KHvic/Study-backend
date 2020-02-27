package app

import (
	"github.com/KHvic/quiz-backend/pkg/constant"
	"github.com/gin-gonic/gin"
)

// Gin ...
type Gin struct {
	C *gin.Context
}

// Response ...
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  constant.GetMsg(errCode),
		Data: data,
	})
	return
}
