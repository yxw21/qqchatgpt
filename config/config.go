package config

import (
	"github.com/yxw21/chatgpt"
	session "github.com/yxw21/chatgpt/session"
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
}

var (
	Instance *Config
	Session  *session.Session
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
	}
	if Instance.AIUsername != "" && Instance.AIPassword != "" && Instance.Key != "" {
		Session = session.NewSessionWithCredential(Instance.AIUsername, Instance.AIPassword, Instance.Key).AutoRefresh()
	} else {
		Session = session.NewSessionWithAccessToken(Instance.AccessToken).AutoRefresh()
	}
	qq, err := strconv.ParseInt(os.Getenv("QQ_UIN"), 10, 64)
	if err == nil {
		Instance.QQ = qq
	}
}
