package main

import (
	_ "github.com/razeencheng/demo-go/swaggo-gin/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"zeropiece/Initialize"
	"zeropiece/UtilityMethodsCollection/douyin"
	"zeropiece/common"
)

// @title ONE-Piece
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://server.shkqg.com/

// @contact.name API Support
// @contact.url http://www.shkqg.com/
// @contact.email rakeecao@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8888
// @BasePath /api
func main() {
	go douyin.CheckAllUsersWorkCycle()
	defer common.CACHE.Close()
	defer common.DB.Close()

	// 初始化路由
	routers := initialize.InitRouters()
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
