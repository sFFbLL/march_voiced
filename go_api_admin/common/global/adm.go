package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v7"
	"github.com/olivere/elastic/v7"
	"github.com/silenceper/wechat/v2/officialaccount"
	"gorm.io/gorm"
)

// orm对外全局变量
var Eloquent *gorm.DB

// redis对外全局变量
var Rdb *redis.Client

//Casbin对外全局边变量
var CasbinEnforcer *casbin.SyncedEnforcer

//weixin全局操作变量
var Wx *officialaccount.OfficialAccount

//elasticSearch对外得到全局变量
var ElasticSearch *elastic.Client
