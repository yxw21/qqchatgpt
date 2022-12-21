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
	SessionToken    string
	FriendAddPolicy string
	TokenFile       string
	DeviceFile      string
}

var Instance *Config

var Chats = make(map[int64]*chatgpt.Chat)

func init() {
	Instance = &Config{
		Password:        os.Getenv("QQ_PASSWORD"),
		AIUsername:      os.Getenv("QQ_CHAT_GPT_USERNAME"),
		AIPassword:      os.Getenv("QQ_CHAT_GPT_PASSWORD"),
		Key:             os.Getenv("QQ_KEY"),
		SessionToken:    os.Getenv("QQ_CHAT_GPT_TOKEN"),
		FriendAddPolicy: os.Getenv("QQ_CHAT_GPT_POLICY"),
		TokenFile:       "qq.token",
		DeviceFile:      "device.json",
	}
	qq, err := strconv.ParseInt(os.Getenv("QQ_UIN"), 10, 64)
	if err == nil {
		Instance.QQ = qq
	}
}
