package wcy_health_check

import (
	"testing"
	"wcy_health_check/tool"
	"wcy_health_check/work"
)

//测试打卡
func TestDaKaWork(t *testing.T) {
	User := tool.GetCfg().User
	work.SendDaKaPost(User)
}
