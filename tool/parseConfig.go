package tool

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type UserConfig struct {
	StuId        string `yaml:"xh"json:"xh"`
	Openid        string `yaml:"openid"json:"openid"`
	Timestamp     int64  `yaml:"timestamp"json:"timestamp"`
	Name          string `yaml:"name"json:"name"`
	Gender        string `yaml:"xb"json:"xb"`
	LocationBig   string `yaml:"locationBig"json:"locationbig"`
	LocationSmall string `yaml:"locationSmall"json:"locationsmall"`
	Latitude      string `yaml:"latitude" json:"latitude"`
	Longitude     string `yaml:"longitude" json:"longitude"`
	Szdq          string `yaml:"szdq" json:"szdq"`
	Xxdz          string `yaml:"xxdz" json:"xxdz"`
	Ywjcqzbl      string `yaml:"ywjcqzbl"json:"ywjcqzbl"`
	Ywjchblj      string `yaml:"ywjchblj" json:"ywjchblj"`
	Xjzdywqzb     string `yaml:"xjzdywqzb" json:"xjzdywqzb"`
	Twsfzc        string `yaml:"twsfzc" json:"twsfzc"`
	Ywytdzz       string `yaml:"ywytdzz"json:"ywytdzz"`
	Remarks       string `yaml:"beizhu" json:"remarks"`
	Mrdkkey       string `yaml:"mrdkkey"json:"mrdkkey"`
	Jkmresult     string `yaml:"jkmresult"json:"jkmresult"`
}


type TimerConfig struct {
     Hour  int  `yaml:"hour"json:"hour"`
	 Minute int  `yaml:"minute"json:"minute"`
}
type CreateOpenId struct {
	Openid          int  `yaml:"openid"json:"openid"`
	UnifiedAuthCode int  `yaml:"yktid"json:"yktid"`
	Password   string `yaml:"passwd"json:"passwd"`
	Timestamp  int64 `yaml:"timestamp"json:"timestamp"`
}
type Config struct {
	User  UserConfig       `yaml:"user"json:"user"`
	TimerClock TimerConfig  `yaml:"timer_clock"json:"timer_clock"`
	CreateOpenId CreateOpenId `yaml:"setting_openid"json:"-"`
}


var cfg  *Config

func GetCfg()*Config{
	return cfg
}

func init()  {
	file,err :=os.Open("./config/app.yaml")
	if err!=nil {
		panic(err)
	}

	bytes,err :=ioutil.ReadAll(file)
	if err !=nil {
		panic(err)
	}
	err =yaml.Unmarshal(bytes,&cfg)
	if err != nil {
		panic(err)
	}
}

