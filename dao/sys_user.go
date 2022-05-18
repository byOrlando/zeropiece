package dao

import "time"

type PeekabooUser struct {
	ID        int       `gorm:"primary_key" json:"id"`
	UserName  string    `json:"username" gorm:"unique_index;"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	AvatarUrl string    `json:"avatar_url"`
	RoleId    int       `json:"role_id"`
	Status    int       `json:"status" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Dept      int       `json:"dept_id" `
	Phone     string    `json:"phone" `
	Email     string    `json:"email" `
	Remark    string    `json:"remark" `
}

type Role struct {
	ID    string `gorm:"primary_key" json:"id"`
	Name  string `gorm:"unique_index;" json:"name"`
	Value string `json:"value"`
}
