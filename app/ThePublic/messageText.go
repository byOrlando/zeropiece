package ThePublic

import "github.com/silenceper/wechat/v2/officialaccount/message"

func MessageText(msg *message.MixMessage) *message.Reply {
	text := message.NewText(msg.Content)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}

func MessageInside(msg *message.MixMessage) *message.Reply {
	text := message.NewText(msg.Content)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}

func MessageOutSide(msg *message.MixMessage) *message.Reply {
	text := message.NewText(msg.Content)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}
