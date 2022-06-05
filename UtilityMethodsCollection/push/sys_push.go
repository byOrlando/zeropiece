package push

import "time"

func SystemPush(title string) {
	var option = InitBasOption()
	times := time.Now().Format("2006-01-02 15:04")
	option.SetTitle(title)
	option.SetGroup("服务器")
	option.SetIcon(IconOption.Get(7))
	option.SetSound(SoundOption.Get(4))
	option.SetText(times)
	option.Push()
}
