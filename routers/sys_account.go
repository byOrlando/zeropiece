package routers

import (
	"github.com/gin-gonic/gin"
	"zeropiece/app/account"
	"zeropiece/middleware"
)

func InitAccountRouter(Router *gin.RouterGroup) {
	UserGroup := Router.Group("/account").Use(middleware.JWTAuth())
	// 获取账户信息
	UserGroup.GET("/getAccountInfo", account.GetAccount)
}
