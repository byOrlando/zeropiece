package push

import (
	"fmt"
	"zeropiece/UtilityMethodsCollection/requests"
	"zeropiece/common"
)

type PushOption struct {
	Text      string `default:"有新作品发布"`
	AwameID   string `default:"7092771219681873166"`
	Sound     string `default:"birdsong"`
	Icon      string `default:"0"`
	GroupName string `default:"抖音通知"`
	Title     string `default:"哆啦戏猫"`
}

func (Option *PushOption) Init() {
	Option.Text = "有新作品发布"
	Option.AwameID = "7092771219681873166"
	Option.Sound = "birdsong"
	Option.Icon = "https://shkqg.com/bark/0.png"
	Option.GroupName = "抖音通知"
	Option.Title = "哆啦戏猫"
}

func (Option *PushOption) Push() {
	client := requests.ClientOption{}
	token := common.CONFIG.Push.PushIphoneToken
	client.Url = fmt.Sprintf("https://bark.shkqg.com/%s/%s/%s", token, Option.Title, Option.Text)
	client.Headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
	}
	client.Params = map[string]string{
		"url":   fmt.Sprintf("snssdk1128://aweme/detail/%s?refer=web&gd_label=click_wap_profile_feature&appParam=&needlaunchlog=1", Option.AwameID),
		"sound": Option.Sound,
		"icon":  Option.Icon,
		"group": Option.GroupName,
	}
	str, err := client.Get()
	if err != nil {
		common.LOG.Error(err.Error())
		return
	}
	common.LOG.Info(str)
}
