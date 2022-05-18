package middleware

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func ResponseSucc(c *gin.Context, msg string, data interface{}) {
	c.JSON(0, Response{Code: 200, Message: msg, Result: data, Type: "success"})
	return
}

func ResponseFail(c *gin.Context, code int, msg string) {
	c.JSON(code, Response{Code: code, Message: msg, Result: nil, Type: "error"})
	return
}
