package server

import (
	"go.uber.org/zap"
	initialize "zeropiece/Initialize"
	"zeropiece/UtilityMethodsCollection/tasks"
	"zeropiece/common"
)

func Run() {
	defer common.CACHE.Close()
	defer common.DB.Close()
	// 启动轮训任务
	go tasks.Run()
	// 初始化路由
	routers := initialize.InitRouters()
	err := routers.Run(":" + common.CONFIG.System.Port)
	if err != nil {
		common.LOG.Error("routers run error", zap.Error(err))
	}
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
