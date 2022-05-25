package admin

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zeropiece/middleware"
	"zeropiece/service"
	"zeropiece/structData"
)

func Login(c *gin.Context) {
	// 登陆逻辑
	var loginData structData.LoginData
	if err := c.ShouldBind(&loginData); err != nil {
		middleware.ResponseFail(c, 201, err.Error())
	}
	user, msg, isPass := service.LoginCheck(loginData.Username, loginData.Password)
	if !isPass {
		middleware.ResponseFail(c, 201, msg)
	} else {
		var res structData.Claims
		res.UserID = strconv.Itoa(user.ID)
		res.UserName = user.UserName
		res.Status = user.Status
		res.RoleId = user.RoleId
		token, msg, ok := middleware.GenerateToken(&res)
		if !ok {
			middleware.ResponseFail(c, 201, msg)
		} else {
			middleware.ResponseSucc(c, msg, structData.LoginDataOut{Token: token})
		}

	}

}

func LogOut(c *gin.Context) {
	// 退出逻辑
	middleware.ResponseSucc(c, "退出成功", nil)
}
