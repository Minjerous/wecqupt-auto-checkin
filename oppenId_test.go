package wcy_health_check

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"testing"
	"time"
	"wcy_health_check/consts"
	"wcy_health_check/tool"
	"wcy_health_check/work"
)
func TestSetOpenID(t *testing.T) {
	OpenidCfg:=tool.GetCfg().CreateOpenId
	OpenidCfg.Timestamp = time.Now().Unix()
	fmt.Println(OpenidCfg)
	b, err := json.Marshal(OpenidCfg)
	if err != nil {
		log.Println(err)
		return
	}
	decodeString := base64.StdEncoding.EncodeToString(b)
	ReqBodyStruct := work.Request{Key: decodeString}
	B,err:=json.Marshal(ReqBodyStruct)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := req.C().EnableDumpAllWithoutResponse().EnableDebugLog()
	Res, err := client.R().EnableDump().SetBodyJsonBytes(B).Post(consts.UrlPostOpenId)
	if err != nil {
		log.Println(err)
		return
	}
	if Res.IsSuccess() {
		fmt.Println(Res)
		fmt.Println("success")
	}else {
		fmt.Println(Res)
	}
}