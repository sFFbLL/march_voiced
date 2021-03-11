package wx

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"project/common/global"
)

func Init() error {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     "wx58f1a0720acc69f4",
		AppSecret: "24fed072517ffb84a7df0343aaeb6ada",
		Token:     "wodeshijie",
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	//获取微信实例
	global.Wx = wc.GetOfficialAccount(cfg)
	return nil
}
