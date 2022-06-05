package dao

import (
	"github.com/jinzhu/gorm"
	"zeropiece/impl/wechat"
)

type InsideWeChatOpenID struct {
	gorm.Model
	OpenID     string             `gorm:"type:varchar(100);unique_index;" json:"open_id"`
	WorkNumber string             `gorm:"type:varchar(10);unique_index;" json:"work_number"`
	Status     string             `gorm:"type:varchar(1);default:'1'" json:"status"`
	AuthCode   wechat.Permissions `gorm:"default:9" json:"auth_code"`
}
