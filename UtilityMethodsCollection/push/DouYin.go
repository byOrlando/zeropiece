package push

import "fmt"

func DouYinPush(videoId string, title string, text string) {
	var option = InitBasOption()
	option.SetTitle(title)
	option.SetGroup("抖音新作品")
	option.SetIcon(IconOption.Get(0))
	option.SetSound(SoundOption.Get(0))
	option.SetText(text)
	option.SetLevel(LevelOption.Get(0))
	option.SetUrl(fmt.Sprintf("snssdk1128://aweme/detail/%s?refer=web&gd_label=click_wap_profile_feature&appParam=&needlaunchlog=1", videoId))
	option.Push()
}
