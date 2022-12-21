package event

import (
	"github.com/Mrs4s/MiraiGo/client"
	"net/http"
	"net/url"
	"qqgpt/config"
	"regexp"
	"strings"
)

func NewFriendRequest(client *client.QQClient, event *client.NewFriendRequest) {
	content := event.Message
	if config.Instance.FriendAddPolicy == "ignore" {
		return
	}
	if config.Instance.FriendAddPolicy == "agree" {
		event.Accept()
		return
	}
	if config.Instance.FriendAddPolicy == "reject" {
		event.Reject()
		return
	}
	sli := strings.Split(config.Instance.FriendAddPolicy, ",")
	if len(sli) < 2 {
		return
	}
	verify := false
	action := sli[0]
	filter := strings.TrimSpace(sli[1])
	if strings.HasPrefix(filter, "https://") || strings.HasPrefix(filter, "http://") {
		// remote verify
		resp, _ := http.Get(filter + "/" + url.PathEscape(content))
		if resp != nil {
			verify = resp.StatusCode == 201
			_ = resp.Body.Close()
		}
	} else {
		// regexp
		reg, _ := regexp.Compile(filter)
		if reg != nil {
			verify = reg.Match([]byte(content))
		}
	}
	if action == "agree" && verify {
		event.Accept()
	}
}
