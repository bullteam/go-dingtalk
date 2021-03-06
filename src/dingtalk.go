package dingtalk

import (
	"net/http"
	"sync"
	"time"
)

/*
*	date: 2018/05/20
	modify date: 2019/11/2
*	version: 0.2
*	author: xiangwenwen(icepy)
*   modify author: wutongci
 */

type DingTalkClient struct {
	DTConfig              *DTConfig
	TopConfig             *TopConfig
	HTTPClient            *http.Client
	AccessToken           string
	SSOAccessToken        string
	SNSAccessToken        string
	SuiteAccessToken      string
	AccessTokenCache      Cache
	TicketCache           Cache
	SSOAccessTokenCache   Cache
	SNSAccessTokenCache   Cache
	SuiteAccessTokenCache Cache
	DevType               string
	Locker                *sync.Mutex
}

type TopConfig struct {
	TopFormat     string // json xml byte
	TopV          string
	TopSimplify   bool
	TopSecret     string
	TopSignMethod string
}

type DTConfig struct {
	TopConfig
	AppKey		  string
	AppSecret     string
	CorpID        string
	CorpSecret    string
	AgentID       string
	SuiteKey      string
	SuiteSecret   string
	SuiteTicket   string
	ChannelSecret string
	SSOSecret     string
	SNSAppID      string
	SNSSecret     string
	CachePath 	  string
}

type DTIsvGetCompanyInfo struct {
	AuthCorpID      string
	PermanentCode   string
	AuthAccessToken string
}

func NewDingTalkClient(devType string, config *DTConfig) *DingTalkClient {
	c := &DingTalkClient{
		DTConfig: &DTConfig{},
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		TopConfig: &TopConfig{
			TopFormat:     topFormat,
			TopSecret:     topSecret,
			TopSignMethod: topSignMethod,
			TopSimplify:   topSimplify,
			TopV:          topV,
		},
		AccessTokenCache:      NewFileCache(config.CachePath + "." + devType + "_access_token_file"),
		TicketCache:           NewFileCache(config.CachePath + "." + devType + "_ticket_file"),
		SSOAccessTokenCache:   NewFileCache(config.CachePath + "." + devType + "_sso_acess_token_file"),
		SNSAccessTokenCache:   NewFileCache(config.CachePath + "." + devType + "_sns_access_token_file"),
		SuiteAccessTokenCache: NewFileCache(config.CachePath + "." + devType + "_suite_access_token_file"),
		Locker:                new(sync.Mutex),
		DevType:               devType,
	}
	if config != nil {
		if config.TopFormat != "" {
			c.TopConfig.TopFormat = config.TopFormat
		}
		if config.TopV != "" {
			c.TopConfig.TopV = config.TopV
		}
		if config.TopSecret != "" {
			c.TopConfig.TopSecret = config.TopSecret
		}
		if config.TopSignMethod != "" {
			c.TopConfig.TopSignMethod = config.TopSignMethod
		}
		if config.TopSimplify {
			c.TopConfig.TopSimplify = config.TopSimplify
		}
		c.DTConfig.AppKey = config.AppKey
		c.DTConfig.AppSecret = config.AppSecret
		c.DTConfig.CorpID = config.CorpID
		c.DTConfig.AgentID = config.AgentID
		c.DTConfig.CorpSecret = config.CorpSecret
		c.DTConfig.SSOSecret = config.SSOSecret
		c.DTConfig.ChannelSecret = config.ChannelSecret
		c.DTConfig.SNSAppID = config.SNSAppID
		c.DTConfig.SNSSecret = config.SNSSecret
		c.DTConfig.SuiteKey = config.SuiteKey
		c.DTConfig.SuiteSecret = config.SuiteSecret
		c.DTConfig.SuiteTicket = config.SuiteTicket
	}
	return c
}

func NewDingTalkISVClient(config *DTConfig) *DingTalkClient {
	return NewDingTalkClient("isv", config)
}

func NewDingTalkCompanyClient(config *DTConfig) *DingTalkClient {
	return NewDingTalkClient("company", config)
}

func NewDingTalkMiniClient(config *DTConfig) *DingTalkClient {
	return NewDingTalkClient("personalMini", config)
}
