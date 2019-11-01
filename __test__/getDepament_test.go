package dingtalkTest

import (
	"testing"
)

func Test_GetDepartmentList(t *testing.T) {
	c := GetCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
	list, err := c.DepartmentList(1,"zh_CN")
	if err != nil {
		t.Error("测试未能获取部门列表")
	} else {
		t.Log("测试获取部门列表通过", list)
	}
}
