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
	MsgRetry        int
}

var (
	Instance *Config
	Session  *chatgpt.Session
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
		MsgRetry:        3,
	}
	if Instance.AIUsername != "" && Instance.AIPassword != "" && Instance.Key != "" {
		Session = chatgpt.NewSessionWithCredential(Instance.AIUsername, Instance.AIPassword, Instance.Key).AutoRefresh()
	} else {
		Session = chatgpt.NewSessionWithAccessToken(Instance.AccessToken).AutoRefresh()
	}
	qq, err := strconv.ParseInt(os.Getenv("QQ_UIN"), 10, 64)
	if err == nil {
		Instance.QQ = qq
	}
	msgRetry, err := strconv.Atoi(os.Getenv("QQ_MSG_RETRY"))
	if err == nil && msgRetry > 0 {
		Instance.MsgRetry = msgRetry
	}
}
