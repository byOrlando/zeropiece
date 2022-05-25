package tasks

import "zeropiece/common"

func Run() {
	defer common.Error()()
	tasks := RegisteredTask()
	for _, task := range tasks {
		go task()
	}
}
