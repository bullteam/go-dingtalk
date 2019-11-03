package dingtalkTest

import (
	dingtalk "github.com/bullteam/go-dingtalk/src"
	"testing"
)

func Test_UserInfoByUserId(t *testing.T)  {
	c := GetCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
	companyInfo := dingtalk.DTIsvGetCompanyInfo{
		AuthCorpID:      "",
		PermanentCode:   "",
		AuthAccessToken: "",
	}
	Info, err := c.UserInfoByUserId("0705603767-509393010","zh_CN",&companyInfo)
	if err != nil {
		t.Error("测试未能获取用户信息")
	} else {
		t.Log("测试获取用户信息通过", Info)
	}
}

func Test_UserSimpleList(t *testing.T)  {
	c := GetCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
	List, err := c.UserSimpleList(142341752)
	if err != nil {
		t.Error("测试未能获取用户列表")
	} else {
		t.Log("测试获取用户列表通过", List)
	}
}
