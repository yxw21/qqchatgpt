package event

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/yxw21/chatgpt"
	"qqgpt/config"
	"qqgpt/helpers"
)

func PrivateMessage(client *client.QQClient, event *message.PrivateMessage) {
	_, content := helpers.GetTextElementFromElements(client.Uin, event.Elements)
	if content == "" {
		client.SendPrivateMessage(event.Sender.Uin, &message.SendingMessage{
			Elements: []message.IMessageElement{message.NewText("暂时只支持文本和表情消息")},
		})
		return
	}
	if _, ok := config.Chats[event.Sender.Uin]; !ok {
		config.Chats[event.Sender.Uin] = chatgpt.NewChat(config.Session)
	}
	res, err := config.Chats[event.Sender.Uin].Send(content)
	if err != nil {
		client.SendPrivateMessage(event.Sender.Uin, &message.SendingMessage{
			Elements: []message.IMessageElement{message.NewText(err.Error())},
		})
	} else {
		client.SendPrivateMessage(event.Sender.Uin, &message.SendingMessage{
			Elements: []message.IMessageElement{message.NewText(res.Message.Content.Parts[0])},
		})
	}
}
