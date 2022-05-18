package account

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"os"
	"zeropiece/common"
	"zeropiece/dao"
	"zeropiece/middleware"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	AushUserInfo := middleware.GetClaims(c)
	var User dao.PeekabooUser
	if err := common.DB.Debug().Where("id = ?", AushUserInfo.UserID).First(&User).Error; err != nil {
		middleware.ResponseFail(c, 401, "获取用户信息失败")
		return
	}
	// 删除原来的头像
	if User.AvatarUrl != "" {
		err := os.Remove(User.AvatarUrl)
		if err != nil {
			os.Remove(string('.') + User.AvatarUrl)
		}
	}
	fileName := "/static/avatar/" + uuid.NewV4().String() + ".png"
	User.AvatarUrl = fileName
	if err := common.DB.Debug().Save(&User).Error; err != nil {
		middleware.ResponseFail(c, 401, "保存头像失败")
		return
	}
	// 上传文件到指定的路径
	err := c.SaveUploadedFile(file, "."+fileName)
	if err != nil {
		return
	}
	middleware.ResponseSucc(c, "上传成功", fileName)
}
