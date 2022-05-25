package admin

import (
	"github.com/gin-gonic/gin"
	"time"
	"zeropiece/common"
	"zeropiece/common/utils"
	"zeropiece/dao"
	"zeropiece/middleware"
	"zeropiece/structData"
)

func Register(c *gin.Context) {

	// 参数校验
	var params structData.RegisterUserData

	if err := c.ShouldBindJSON(&params); err != nil {
		middleware.ResponseFail(c, 201, "非法数据")
		return
	}

	err := common.DB.Debug().Where("user_name = ?", params.Account).Find(&dao.PeekabooUser{}).Error
	if err == nil {
		middleware.ResponseFail(c, 201, "用户已存在")
		return
	}
	err = common.DB.Debug().Where("phone = ?", params.Mobile).Find(&dao.PeekabooUser{}).Error
	if err == nil {
		middleware.ResponseFail(c, 201, "手机号码已存在")
		return
	}
	user := dao.PeekabooUser{
		UserName:  params.Account,
		Nickname:  params.Account,
		Password:  utils.Md5(params.Password + common.PWsalt),
		AvatarUrl: "",
		RoleId:    56,
		Status:    1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Dept:      0,
		Phone:     params.Mobile,
		Email:     "",
		Remark:    "",
	}
	err = common.DB.Debug().Create(&user).Error
	if err != nil {
		middleware.ResponseFail(c, 201, "注册失败")
		return
	}
	middleware.ResponseSucc(c, "注册成功", map[string]string{"username": "12312", "msg": "注册成功"})

	//middleware.ResponseSucc(c, "注册成功", map[string]string{"username": "123456"})

}
