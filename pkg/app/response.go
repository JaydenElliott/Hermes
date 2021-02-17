package app

import (
	"github.com/gin-gonic/gin"

	"arcstack/arcstack-chat-server/pkg/api_response"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  api_response.GetMsg(errCode),
		Data: data,
	})
	return
}
