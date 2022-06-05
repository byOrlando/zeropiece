package ThePublic

import (
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

func MessageVoice(msg *message.MixMessage) *message.Reply {
	voice := message.NewVoice(msg.MediaID)
	fmt.Println(msg.Content)
	return &message.Reply{MsgType: message.MsgTypeVoice, MsgData: voice}
}
