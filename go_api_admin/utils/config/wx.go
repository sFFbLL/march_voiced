package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Wx struct {
	AppID     string
	AppSecret string
	Token     string
}

// InitWx 初始化Wx配置
func InitWx(cfg *viper.Viper) *Wx {
	fmt.Println(cfg.GetString("appid"), "111111")
	wx := &Wx{
		AppID:     cfg.GetString("appid"),
		AppSecret: cfg.GetString("appsecret"),
		Token:     cfg.GetString("token"),
	}
	return wx
}

var WxConfig = new(Wx)
