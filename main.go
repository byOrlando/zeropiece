package main

import (
	"zeropiece/Initialize"
	"zeropiece/UtilityMethodsCollection/douyin"
	"zeropiece/common"
)

func main() {
	go douyin.CheckAllUsersWorkCycle()
	defer common.CACHE.Close()
	defer common.DB.Close()

	// 初始化路由
	routers := initialize.InitRouters()
	_ = routers.Run(":" + common.CONFIG.System.Port)
}

func init() {
	//初始华配置
	initialize.InitConf()
	// 初始化日志
	initialize.InitLog()
	//初始化redis
	initialize.InitCache()
	//初始化数据库
	initialize.InitDb()
}
