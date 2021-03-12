package wx

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"project/common/global"
	"project/utils/config"
)

func Init(cfg *config.Wx) error {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	wxCfg := &offConfig.Config{
		AppID:     cfg.AppID,
		AppSecret: cfg.AppSecret,
		Token:     cfg.Token,
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	//获取微信实例
	global.Wx = wc.GetOfficialAccount(wxCfg)
	return nil
}
