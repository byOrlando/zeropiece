package wechat

type Permissions int

const (
	NonePermissions Permissions = 0
	Manager         Permissions = 3
	Customer        Permissions = 2
	User            Permissions = 1
	Admin           Permissions = 9
)
