package admin

import (
	"github.com/gin-gonic/gin"
	"zeropiece/common/utils"
	"zeropiece/middleware"
)

func GetIp(c *gin.Context) {
	// 获取客户端IP
	address := utils.GetLocation(c.ClientIP())
	middleware.ResponseSucc(c, "获取IP地址成功", address)
}
