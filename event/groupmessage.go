package event

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/yxw21/chatgpt"
	"qqgpt/config"
	"qqgpt/helpers"
)

func GroupMessage(client *client.QQClient, event *message.GroupMessage) {
	isAt, content := helpers.GetTextElementFromElements(client.Uin, event.Elements)
	if isAt {
		if _, ok := config.Chats[event.Sender.Uin]; !ok {
			config.Chats[event.Sender.Uin] = chatgpt.NewChat(config.Session)
		}
		for i := 0; i < config.Instance.MsgRetry; i++ {
			res, err := config.Chats[event.Sender.Uin].Send(content)
			if err == nil {
				client.SendGroupMessage(event.GroupCode, &message.SendingMessage{
					Elements: []message.IMessageElement{message.NewAt(event.Sender.Uin), message.NewText(res.Message.Content.Parts[0])},
				})
				break
			}
			if i == config.Instance.MsgRetry-1 {
				client.SendGroupMessage(event.GroupCode, &message.SendingMessage{
					Elements: []message.IMessageElement{message.NewAt(event.Sender.Uin), message.NewText(err.Error())},
				})
			}
		}
	}
}
