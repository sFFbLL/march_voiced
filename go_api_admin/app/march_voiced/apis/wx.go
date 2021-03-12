package apis

import (
	"fmt"
	"time"

	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/service"
	"project/common/global"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
	"go.uber.org/zap"
)

func WxGetTicket(c *gin.Context) {
	server := global.Wx.GetServer(c.Request, c.Writer)
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//TODO 对接收到的消息以及处理
		text := message.NewText("欢迎来到三月之声")
		if msg.Event == "SCAN" {
			text = message.NewText("欢迎来到三月之声")
			_, err := global.Rdb.Set(fmt.Sprintf("%v", msg.EventKey), fmt.Sprintf("%v", msg.FromUserName), 300*time.Second).Result()
			if err != nil {
				zap.L().Error("Wx scan message failed", zap.Error(err))
				return &message.Reply{MsgType: message.MsgTypeText}
			}
		}
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})
	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		zap.L().Error("Wx message server failed", zap.Error(err))
		return
	}
	server.Send()
}

func CheckIsSucess(c *gin.Context) {
	str := c.Query("strData")
	result, err := global.Rdb.Get(str).Result()
	if err != nil {
		zap.L().Error("GetAccessToken failed", zap.Error(err))
		app.ResponseError(c, app.CodeWxTickerFail)
		return
	}
	user := global.Wx.GetUser()
	if result == "" {
		resultData := bo.MarchLoginData{
			OpenId: "",
			Token:  "",
			Status: 0,
		}
		app.ResponseSuccess(c, resultData)
		return
	}
	info, err := user.GetUserInfo(result)
	if err != nil {
		zap.L().Error("GetAccessToken failed", zap.Error(err))
		app.ResponseError(c, app.CodeWxOuttime)
		return
	}
	oauthInfo := oauth.UserInfo{
		OpenID:     info.OpenID,
		Nickname:   info.Nickname,
		Sex:        info.Sex,
		Province:   info.Province,
		City:       info.City,
		Country:    info.Country,
		HeadImgURL: info.Headimgurl,
		Unionid:    info.UnionID,
	}
	login := new(service.Login)
	data, err := login.LoginService(c, oauthInfo)
	if err != nil {
		zap.L().Error("CreatDaoWeixinUserInfo failed", zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}
	// 如果是statu 是1则请求补全信息接口，if是2则获取到token登录成功
	app.ResponseSuccess(c, data)
}

func SignInfo(c *gin.Context) {
	url := c.Query("url")
	// 获取js操作对象信息
	js := global.Wx.GetJs()
	//获取js签名
	config, err := js.GetConfig(url)
	if err != nil {
		zap.L().Error("get wx sign failed", zap.Error(err))
		app.ResponseError(c, app.CodeWxGzhSignFail)
	}
	app.ResponseSuccess(c, config)
}
