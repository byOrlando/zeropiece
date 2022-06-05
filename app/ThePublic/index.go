package ThePublic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"zeropiece/common"
)

func ServeWechat(c *gin.Context) {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:          common.CONFIG.Wechat.AppID,
		AppSecret:      common.CONFIG.Wechat.AppSecret,
		Token:          common.CONFIG.Wechat.Token,
		EncodingAESKey: common.CONFIG.Wechat.AesKey,
		Cache:          memory,
	}
	officialAccount := wc.GetOfficialAccount(cfg)

	// 传入request和responseWriter
	server := officialAccount.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(MessageHandler)

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}
