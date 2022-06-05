package ThePublic

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

func MessageHandler(msg *message.MixMessage) *message.Reply {
	//TODO
	//回复消息：演示回复用户发送的消息
	switch msg.MsgType {
	case message.MsgTypeText:
		return MessageText(msg)
	case message.MsgTypeImage:
		return MessageImage(msg)
	case message.MsgTypeVoice:
		return MessageVoice(msg)
	case message.MsgTypeEvent:
		return MessageEvent(msg)

	default:
		return nil
	}
}
