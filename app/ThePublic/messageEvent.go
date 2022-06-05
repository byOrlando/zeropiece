package ThePublic

import "github.com/silenceper/wechat/v2/officialaccount/message"

func MessageEvent(msg *message.MixMessage) *message.Reply {
	switch msg.Event {
	case message.EventClick:
		return MessageClick(msg)
	case message.EventView:
		return MessageView(msg)
	case message.EventSubscribe:
		return MessageSubscribe(msg)
	case message.EventUnsubscribe:
		return MessageUnsubscribe(msg)

	default:
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: "点击类型错误"}
	}

}
func MessageClick(msg *message.MixMessage) *message.Reply {
	// TODO 可以根据msg.EventKey判断点击的是哪个菜单
	text := message.NewText(msg.EventKey)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}

func MessageView(msg *message.MixMessage) *message.Reply {
	// TODO 可以根据msg.EventKey判断点击的是哪个菜单
	text := message.NewText(msg.EventKey)
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}

func MessageSubscribe(msg *message.MixMessage) *message.Reply {
	// 关注公众号事件
	text := message.NewText("感谢您的关注！")
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}

func MessageUnsubscribe(msg *message.MixMessage) *message.Reply {
	// 取消关注公众号事件
	text := message.NewText("取消关注")
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
}
