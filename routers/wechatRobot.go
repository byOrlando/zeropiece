package routers

import (
	"github.com/gin-gonic/gin"
	"zeropiece/app/ThePublic"
	"zeropiece/controller/wechat"
)

func InitWechatRobot(Router *gin.RouterGroup) {
	Router.POST("", ThePublic.ServeWechat)
	Router.GET("", ThePublic.ServeWechat)
}

func InitWechat(Router *gin.RouterGroup) {
	WechatGroup := Router.Group("/wechat")
	WechatGroup.GET("/OpenID", wechat.OpenIdGet)
	WechatGroup.POST("/OpenID", wechat.OpenIdPost)
	WechatGroup.DELETE("/OpenID", wechat.OpenIdDelete)
	WechatGroup.PUT("/OpenID", wechat.OpenIdPut)
}
