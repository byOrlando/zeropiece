package initialize

import (
	"zeropiece/UtilityMethodsCollection/wsocket"
	"zeropiece/common"

	"net/http"
	"zeropiece/middleware"
	"zeropiece/routers"

	"github.com/gin-gonic/gin"
)

func InitRouters() (Router *gin.Engine) {
	gin.SetMode(common.CONFIG.Project.Mode)
	Router = gin.Default()

	Router.Use(middleware.Cors())
	common.LOG.Info("use middleware cors logger recovery")
	Router.Static("/bark", "../onepiece/dist")
	//Router.LoadHTMLGlob("templates/*")
	Router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/bark/index.html")
	})
	Router.GET("/ws", wsocket.WsHandler)

	// api group, use middleware Logger to log the ip and time of access
	ApiV1Group := Router.Group("/api")
	routers.InitBaseRouter(ApiV1Group)
	routers.InitAccountRouter(ApiV1Group)
	routers.InitDouYinRouter(ApiV1Group)
	routers.InitWechat(ApiV1Group)
	// WeChat group, use middleware Logger to log the ip and time of access
	WechatRobotGroup := Router.Group("/NewChat")
	routers.InitWechatRobot(WechatRobotGroup)
	common.LOG.Info("routers register success")

	return
}
