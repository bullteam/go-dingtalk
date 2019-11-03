package dingtalkTest

import (
	"os"
	"github.com/bullteam/go-dingtalk/src"
)

func GetCompanyDingTalkClient() *dingtalk.DingTalkClient {
	config := &dingtalk.DTConfig{
		CorpID:     os.Getenv("CorpId"),
		CorpSecret: os.Getenv("CorpSecret"),
		AgentID:    os.Getenv("AgentID"),
		SSOSecret:  os.Getenv("SSOSecret"),
		SNSAppID:   os.Getenv("SNSAppID"),
		SNSSecret:  os.Getenv("SNSSecret"),
		AppKey:     os.Getenv("AppKey"),
		AppSecret: os.Getenv("AppSecret"),
		CachePath: "",
	}
	c := dingtalk.NewDingTalkCompanyClient(config)
	return c
}
