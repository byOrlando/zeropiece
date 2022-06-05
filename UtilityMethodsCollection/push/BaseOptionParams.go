package push

import (
	"fmt"
	"zeropiece/UtilityMethodsCollection/requests"
	"zeropiece/common"
)

type BaseOptionParams struct {
	title    string `default:"标题"`
	group    string `default:"分组"`
	icon     string `default:"图标"`
	sound    string `default:"提示音"`
	text     string `default:"内容"`
	level    string `default:"中断级别"`
	url      string `default:"跳转地址"`
	badge    string `default:"角标"`
	copy     string `default:"复制内容"`
	autoCopy string `default:"自动复制内容"`
	archive  string `default:"保存消息"`
}

func (s *BaseOptionParams) SetTitle(title string) {
	s.title = title
}
func (s *BaseOptionParams) GetTitle() string {
	return s.title
}
func (s *BaseOptionParams) SetGroup(group string) {
	s.group = group
}
func (s *BaseOptionParams) GetGroup() string {
	return s.group
}
func (s *BaseOptionParams) SetIcon(icon string) {
	s.icon = icon
}
func (s *BaseOptionParams) GetIcon() string {
	return s.icon
}
func (s *BaseOptionParams) SetSound(sound string) {
	s.sound = sound
}

func (s *BaseOptionParams) GetSound() string {
	return s.sound
}

func (s *BaseOptionParams) SetText(text string) {
	s.text = text
}

func (s *BaseOptionParams) GetText() string {
	return s.text
}

func (s *BaseOptionParams) SetLevel(level string) {
	s.level = level
}
func (s *BaseOptionParams) GetLevel() string {
	return s.level
}

func (s *BaseOptionParams) SetUrl(url string) {
	s.url = url
}
func (s *BaseOptionParams) GetUrl() string {
	return s.url
}
func (s *BaseOptionParams) SetBadge(badge string) {
	s.badge = badge
}
func (s *BaseOptionParams) GetBadge() string {
	return s.badge
}
func (s *BaseOptionParams) SetCopy(copy string) {
	s.copy = copy
}
func (s *BaseOptionParams) GetCopy() string {
	return s.copy
}
func (s *BaseOptionParams) SetAutoCopy(autoCopy string) {
	s.autoCopy = autoCopy
}
func (s *BaseOptionParams) GetAutoCopy() string {
	return s.autoCopy
}
func (s *BaseOptionParams) SetArchive(archive string) {
	s.archive = archive
}
func (s *BaseOptionParams) GetArchive() string {
	return s.archive
}

func InitBasOption() *BaseOptionParams {
	var option = BaseOptionParams{}
	option.SetTitle("服务器通知")
	option.SetGroup("服务器")
	option.SetText("服务器发生错误")
	option.SetIcon(IconOption.Get(0))
	option.SetSound(SoundOption.Get(0))
	option.SetLevel(LevelOption.Get(0))
	option.SetUrl("")
	option.SetBadge("")
	option.SetCopy("")
	option.SetAutoCopy("")
	option.SetArchive("")
	return &option
}

func (s *BaseOptionParams) Push() {
	params := make(map[string]string)
	client := requests.ClientOption{}
	token := common.CONFIG.Push.PushIphoneToken
	client.Url = fmt.Sprintf("https://bark.shkqg.com/%s/%s/%s", token, s.GetTitle(), s.GetText())
	client.Headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
	}
	if s.GetGroup() != "" {
		params["group"] = s.GetGroup()
	}
	if s.GetIcon() != "" {
		params["icon"] = s.GetIcon()
	}
	if s.GetSound() != "" {
		params["sound"] = s.GetSound()
	}
	if s.GetLevel() != "" {
		params["level"] = s.GetLevel()
	}
	if s.GetUrl() != "" {
		params["url"] = s.GetUrl()
	}
	if s.GetBadge() != "" {
		params["badge"] = s.GetBadge()
	}
	if s.GetCopy() != "" {
		params["copy"] = s.GetCopy()
	}
	if s.GetAutoCopy() != "" {
		params["autoCopy"] = s.GetAutoCopy()
	}
	if s.GetArchive() != "" {
		params["archive"] = s.GetArchive()
	}

	client.Params = params
	str, err := client.Get()
	if err != nil {
		common.LOG.Error(err.Error())
		return
	}
	common.LOG.Info(str)
}
