package ThePublic

import "github.com/silenceper/wechat/v2/officialaccount/message"

func MessageImage(msg *message.MixMessage) *message.Reply {
	text := message.NewText(msg.PicURL)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}
