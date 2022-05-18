package routers

import (
	"github.com/gin-gonic/gin"
	"zeropiece/app/account"
	v1 "zeropiece/app/admin"
	"zeropiece/middleware"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	// 用户登录
	Router.POST("/login", v1.Login)
	// 注册
	Router.POST("/register", v1.Register)
	// 上传图片
	Router.POST("/upload", middleware.JWTAuth(), account.Upload)
	// 退出
	Router.GET("/logout", middleware.JWTAuth(), v1.LogOut)
	// 获取用户信息
	Router.GET("/getUserInfo", middleware.JWTAuth(), v1.GetUserInfo)
	// 获取天气情况
	Router.GET("/getWeather", v1.GetWeather)
	// 获取用户ip地址
	Router.GET("/getIpAddress", v1.GetIp)
}
