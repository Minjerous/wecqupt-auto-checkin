package work

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"math/rand"
	"time"
	"wcy_health_check/consts"
	"wcy_health_check/tool"
)

type Request struct {
	Key string `json:"key"`
}
//生成时间
func CreateTime() time.Time{
	CfgTimer := tool.GetCfg().TimerClock
	hour:=CfgTimer.Hour
	now:=time.Now()
	minute := CfgTimer.Minute
	sec := rand.Intn(60)
	return time.Date(now.Year(), now.Month(), now.Day(), hour, minute, sec, 0, time.Local).Add(24*time.Hour)
}
func DaKaWork()  {
	CfgUser := tool.GetCfg().User
	err := SendDaKaPost(CfgUser)
	if err != nil {
		log.Println(err)
		return
	}
	t := CreateTime()
	timer := time.NewTimer(t.Sub(time.Now()))
	for {
		<-timer.C
		err := SendDaKaPost(CfgUser)
		if err != nil {
			log.Println("打卡失败")
			log.Println(err)
		}
		// 计算下一个零点
		t = CreateTime()
		timer.Reset(t.Sub(time.Now()))
	}

}

func SendDaKaPost(cfgUser tool.UserConfig)error {
	cfgUser.Timestamp = time.Now().Unix()
	cfgUser.Mrdkkey = tool.GetMirKey(time.Now().Day(), time.Now().Hour())
	b, err := json.Marshal(cfgUser)
	if err != nil {
		log.Println(err)
		return err
	}
	decodeString := base64.StdEncoding.EncodeToString(b)
	ReqBodyStruct := Request{Key: decodeString}
	B,err:=json.Marshal(ReqBodyStruct)
	if err != nil {
		fmt.Println(err)
		return err
	}
	client := req.C().EnableDumpAllWithoutResponse().EnableDebugLog()
	Res, err := client.R().EnableDump().SetBodyJsonBytes(B).Post(consts.UrlPostDaKa)
	if err != nil {
		return err
	}
	if Res.IsSuccess() {
		fmt.Println(Res)
		log.Println("~~~~~打卡成功~~~~~~~~~")
		return err
	}else {
		log.Println("~~~~~打卡失败~~~~~~~~~")
	}
	return err
}
