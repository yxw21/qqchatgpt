package config

import (
	"github.com/yxw21/chatgpt"
	"os"
	"strconv"
)

type Config struct {
	QQ              int64
	Password        string
	AIUsername      string
	AIPassword      string
	Key             string
	AccessToken     string
	FriendAddPolicy string
	TokenFile       string
	DeviceFile      string
	Proxy           string
}

var (
	Instance *Config
	Session  = &chatgpt.Session{}
	Browser  *chatgpt.Browser
	Chats    = make(map[int64]*chatgpt.Chat)
)

func init() {
	Instance = &Config{
		Password:        os.Getenv("QQ_PASSWORD"),
		AIUsername:      os.Getenv("QQ_CHAT_GPT_USERNAME"),
		AIPassword:      os.Getenv("QQ_CHAT_GPT_PASSWORD"),
		Key:             os.Getenv("QQ_KEY"),
		AccessToken:     os.Getenv("QQ_CHAT_GPT_ACCESS_TOKEN"),
		FriendAddPolicy: os.Getenv("QQ_CHAT_GPT_POLICY"),
		TokenFile:       "qq.token",
		DeviceFile:      "device.json",
		Proxy:           os.Getenv("QQ_PROXY"),
	}
	qq, err := strconv.ParseInt(os.Getenv("QQ_UIN"), 10, 64)
	if err == nil && qq > 0 {
		Instance.QQ = qq
	}
	Session.Username = Instance.AIUsername
	Session.Password = Instance.AIPassword
	Session.AccessToken = Instance.AccessToken
}
