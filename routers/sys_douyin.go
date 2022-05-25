package routers

import (
	"github.com/gin-gonic/gin"
	"zeropiece/controller/app_douyin"
	"zeropiece/middleware"
)

func InitDouYinRouter(Router *gin.RouterGroup) {
	DouYinRouter := Router.Group("/douyin").Use(middleware.JWTAuth())
	// 获取抖音用户列表
	DouYinRouter.GET("/getUserListByPage", app_douyin.GetUserList)
	// 设置抖音用户状态
	DouYinRouter.POST("/setUserStatus", app_douyin.SetUserStatus)
	// 更新抖音用户信息
	DouYinRouter.POST("/insetUser", app_douyin.InsetUser)
}
