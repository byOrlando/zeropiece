package app_douyin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zeropiece/common"
	"zeropiece/dao"
	"zeropiece/middleware"
)

type DouYinUserOutList struct {
	Id         int    `json:"id"`
	OrderNo    int    `json:"orderNo"`
	Remark     string `json:"remark"`
	UserName   string `json:"userName"`
	Uid        string `json:"uid"`
	Status     string `json:"status"`
	CreateTime string `json:"createTime"`
}

// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept json
// @Param who query string true "hehe"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /GetUserList [get]

func GetUserList(c *gin.Context) {
	var users []dao.DouyinUser
	var outList []DouYinUserOutList
	common.DB.Find(&users)
	num := 0
	for _, v := range users {
		num += 1
		outList = append(outList, DouYinUserOutList{
			Id:         int(v.ID),
			OrderNo:    num,
			Remark:     v.Remark,
			UserName:   v.UserName,
			Uid:        v.UID,
			Status:     v.Status,
			CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	middleware.ResponseSucc(c, "获取成功", outList)
}

type UserStatusIn struct {
	Status string `json:"status"`
	ID     int    `json:"id"`
}

func SetUserStatus(c *gin.Context) {
	var UserIn UserStatusIn
	if err := c.ShouldBind(&UserIn); err != nil {
		middleware.ResponseFail(c, 201, "参数错误")
		return
	}
	var user dao.DouyinUser
	if err := common.DB.Where("id = ?", UserIn.ID).First(&user).Error; err != nil {
		middleware.ResponseFail(c, 201, "用户不存在")
		return
	}
	user.Status = UserIn.Status
	if err := user.Save(); err != nil {
		middleware.ResponseFail(c, 201, "修改失败")
		return
	}
	middleware.ResponseSucc(c, "设置成功", nil)
}

type DouYinUserIn struct {
	UserName string `json:"userName"`
	Uid      string `json:"uid"`
	Remark   string `json:"remark"`
	Status   string `json:"status"`
}

func InsetUser(c *gin.Context) {
	var inModel DouYinUserIn
	if err := c.ShouldBind(&inModel); err != nil {
		middleware.ResponseFail(c, 201, "参数错误")
		return
	}
	var user dao.DouyinUser
	if err := common.DB.Where("uid = ?", inModel.Uid).First(&user).Error; err == nil {
		fmt.Println("正在执行修改操作")
		user.UserName = inModel.UserName
		user.UID = inModel.Uid
		user.Remark = inModel.Remark
		user.Status = inModel.Status
		if err := user.Save(); err != nil {
			middleware.ResponseFail(c, 201, "修改失败")
			return
		}
		middleware.ResponseSucc(c, "修改成功", nil)
		return

	} else {
		user.UserName = inModel.UserName
		user.UID = inModel.Uid
		user.Remark = inModel.Remark
		user.Status = inModel.Status
		if err := common.DB.Create(&user).Error; err != nil {
			middleware.ResponseFail(c, 201, "插入失败")
			return
		}
		middleware.ResponseSucc(c, "创建成功", nil)
		return
	}

}

func DeleteUser(c *gin.Context) {
	var inModel DouYinUserIn
	if err := c.ShouldBind(&inModel); err != nil {
		middleware.ResponseFail(c, 201, "参数错误")
		return
	}
	var user dao.DouyinUser
	if err := common.DB.Where("uid = ?", inModel.Uid).First(&user).Error; err != nil {
		middleware.ResponseFail(c, 201, "用户不存在")
		return
	}
	if err := common.DB.Delete(&user).Error; err != nil {
		middleware.ResponseFail(c, 201, "删除失败")
		return
	}
	middleware.ResponseSucc(c, "删除成功", nil)
}
