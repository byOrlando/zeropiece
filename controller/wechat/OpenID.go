package wechat

import "github.com/gin-gonic/gin"

func OpenIdGet(c *gin.Context) {
	c.JSON(200, gin.H{"token": "数据获取成功"})
}

func OpenIdPost(c *gin.Context) {
	c.JSON(200, gin.H{"token": "数据保存成功"})
}

func OpenIdDelete(c *gin.Context) {
	c.JSON(200, gin.H{"token": "数据删除成功"})
}
func OpenIdPut(c *gin.Context) {
	c.JSON(200, gin.H{"token": "数据修该成功"})
}
