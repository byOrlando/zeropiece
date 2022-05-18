package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"zeropiece/common"
	"zeropiece/dao"
	"zeropiece/middleware"
	"zeropiece/structData"
)

func GetUserInfo(c *gin.Context) {
	// 获取用户信息
	AushUserInfo := middleware.GetClaims(c)
	fmt.Println(AushUserInfo)
	var User dao.PeekabooUser
	if err := common.DB.Debug().Where("id = ?", AushUserInfo.UserID).First(&User).Error; err != nil {
		middleware.ResponseFail(c, 401, "获取用户信息失败")
		return
	}
	// 获取用户信
	var RoleData dao.Role
	if err := common.DB.Debug().Where("id = ?", User.RoleId).First(&RoleData).Error; err != nil {
		fmt.Println(err)
		common.LOG.Info("获取用户角色信息失败")
	}
	middleware.ResponseSucc(c, "获取用户信息成功", structData.UserInfoOut{
		Role:     []map[string]string{{"raleName": RoleData.Name, "value": RoleData.Value}},
		UserId:   strconv.Itoa(User.ID),
		UserName: User.UserName,
		NickName: User.Nickname,
		Avatar:   User.AvatarUrl,
		Desc:     User.Remark,
	},
	)
}
