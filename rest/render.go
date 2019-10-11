package rest

import (
	"github.com/gin-gonic/gin"
)

type RespJsonObj struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func RespJson(c *gin.Context, code int, data interface{}) {
	errMsg := StatusText(code)
	if err, ok := data.(error); ok {
		errMsg = err.Error()
	}
	result := &RespJsonObj{
		Code: code,
		Msg:  errMsg,
		Data: data,
	}
	c.JSON(HttpStatusCode(code), result)
}

func RespJsonWithAbord(c *gin.Context, code int, data interface{}) {
	result := &RespJsonObj{
		Code: code,
		Msg:  StatusText(code),
		Data: data,
	}
	c.AbortWithStatusJSON(HttpStatusCode(code), result)
}
