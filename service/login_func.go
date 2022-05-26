package service

import (
	"zeropiece/common"
	"zeropiece/common/utils"
	"zeropiece/conf"
	"zeropiece/dao"
)

// LoginCheck 登录验证
func LoginCheck(username string, password string) (user *dao.PeekabooUser, msg string, isPass bool) {
	password = utils.Md5(password + conf.PassWordSalt)
	//fmt.Println(password)
	user, ok := CheckUserExist(username, password)

	if !ok {
		return user, "用户或密码不正确", false
	}
	return user, "登录成功", true
}

// CheckUserExist 验证用户是否存在
func CheckUserExist(username string, password string) (user *dao.PeekabooUser, ok bool) {
	user = new(dao.PeekabooUser)
	common.DB.Where(dao.PeekabooUser{UserName: username, Password: password, Status: 1}).Preload("SysRole").First(user)
	return user, user.ID > 0
}
