package tasks

import (
	"fmt"
	"time"
	"zeropiece/UtilityMethodsCollection/douyin"
)

var numQ = 1

func SuperWang() {
	douyin.CheckAllUsersWorkCycle()
	fmt.Printf("The %dth check is complete! -- %s", numQ, time.Now())
	numQ++
}
