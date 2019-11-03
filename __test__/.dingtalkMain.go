package main

import (
	"os"
	"github.com/bullteam/go-dingtalk/src"
)

func main() {
	c := getCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
}

func getCompanyDingTalkClient() *dingtalk.DingTalkClient {
	config := &dingtalk.DTConfig{
		AppKey:     os.Getenv("AppKey"),
		AppSecret: os.Getenv("AppSecret"),
	}
	c := dingtalk.NewDingTalkCompanyClient(config)
	return c
}
