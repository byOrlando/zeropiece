package routers

import (
	"github.com/gin-gonic/gin"
	"zeropiece/controller/account"
	"zeropiece/middleware"
)

func InitAccountRouter(Router *gin.RouterGroup) {
	UserGroup := Router.Group("/account").Use(middleware.JWTAuth())
	// 获取账户信息
	UserGroup.GET("/getAccountInfo", account.GetAccount)
}
