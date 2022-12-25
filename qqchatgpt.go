package main

import (
	"context"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/yxw21/chatgpt"
	"log"
	"qqgpt/config"
	"qqgpt/event"
	"qqgpt/helpers"
)

func main() {
	var (
		qq  *client.QQClient
		err error
	)
	browser, closeBrowser, err := chatgpt.NewBrowser(config.Instance.Key)
	if err != nil {
		log.Fatal(err)
	}
	defer closeBrowser()
	config.Browser = browser
	if config.Instance.AIUsername != "" && config.Instance.AIPassword != "" {
		config.Session = chatgpt.NewSessionWithCredential(browser, config.Instance.AIUsername, config.Instance.AIPassword).AutoRefresh()
	} else {
		config.Session = chatgpt.NewSessionWithAccessToken(browser, config.Instance.AccessToken).AutoRefresh()
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err = helpers.AutoLoadDevice(); err != nil {
		log.Fatal("加载设备信息失败")
	}
	qq, err = helpers.LoginWithToken()
	if err != nil {
		log.Println(err)
	}
	if err != nil && config.Instance.QQ != 0 {
		qq, err = helpers.LoginWithPassword(config.Instance.QQ, config.Instance.Password)
		if err != nil {
			log.Println(err)
		}
	}
	if err != nil {
		qq, err = helpers.LoginWithQRCode(true)
		if err != nil {
			log.Fatal(err)
		}
	}
	qq.PrivateMessageEvent.Subscribe(event.PrivateMessage)
	qq.GroupMessageEvent.Subscribe(event.GroupMessage)
	qq.NewFriendRequestEvent.Subscribe(event.NewFriendRequest)
	<-ctx.Done()
}
