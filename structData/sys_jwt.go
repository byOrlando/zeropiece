package structData

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims 载荷，可以加一些自己需要的信息
type Claims struct {
	UserID   string `json:"user_id"`
	UserName string `json:"username"`
	Status   int    `json:"status"`
	RoleId   int    `json:"roleId"`
	jwt.StandardClaims
}

type UserInfoOut struct {
	Avatar   string              `json:"avatar"`
	Desc     string              `json:"desc"`
	HomePath string              `json:"homePath"`
	Role     []map[string]string `json:"role"`
	UserId   string              `json:"user_Id"`
	UserName string              `json:"username"`
	NickName string              `json:"nickname"`
	Status   int                 `json:"status"`
}
