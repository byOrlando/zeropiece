package account

import (
	"github.com/gin-gonic/gin"
	"zeropiece/middleware"
)

type Account struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	Phone        string `json:"phone"`
	Asddress     string `json:"address"`
}

func GetAccount(c *gin.Context) {
	var account Account
	account.Email = "909038822@qq.com"
	account.Name = "ZeroPiece"
	account.Introduction = "ZeroPiece"
	account.Phone = "18888888888"
	account.Asddress = "ZeroPiece"
	middleware.ResponseSucc(c, "获取用户信息成功", &account)

}
