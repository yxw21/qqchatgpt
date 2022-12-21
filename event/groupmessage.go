package event

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"qqgpt/config"
	"qqgpt/helpers"
)

func GroupMessage(client *client.QQClient, event *message.GroupMessage) {
	isAt, content := helpers.GetTextElementFromElements(client.Uin, event.Elements)
	if isAt {
		if _, ok := config.Chats[event.Sender.Uin]; !ok {
			config.Chats[event.Sender.Uin] = helpers.GetNewChat()
		}
		res, err := config.Chats[event.Sender.Uin].Send(content)
		if err != nil {
			client.SendGroupMessage(event.GroupCode, &message.SendingMessage{
				Elements: []message.IMessageElement{message.NewAt(event.Sender.Uin), message.NewText(err.Error())},
			})
		} else {
			client.SendGroupMessage(event.GroupCode, &message.SendingMessage{
				Elements: []message.IMessageElement{message.NewAt(event.Sender.Uin), message.NewText(res.Message.Content.Parts[0])},
			})
		}
	}
}
