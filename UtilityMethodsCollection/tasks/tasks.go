package tasks

import (
	"zeropiece/UtilityMethodsCollection/douyin"
)

func RegisteredTask() (FuncList []func()) {
	return []func(){
		douyin.CheckAllUsersWorkCycle,
	}
}
