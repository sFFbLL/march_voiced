package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
	"go.uber.org/zap"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/service"
	"project/common/global"
	"project/utils/app"
	"time"
)

func WxGetTicket(c *gin.Context) {
	server := global.Wx.GetServer(c.Request,c.Writer)
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//TODO 对接收到的消息以及处理
		text := message.NewText("111")
		if msg.Event =="SCAN"{
			textOne := message.NewText("欢迎来到三月之声")
			_, err := global.Rdb.Set(fmt.Sprintf("%v",msg.EventKey ),fmt.Sprintf("%v",msg.FromUserName) , 300*time.Second).Result()
			if err != nil {
				return &message.Reply{MsgType: message.MsgTypeText}
			}
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: textOne}
		}
		//text := message.NewText("msg.Content")
		return &message.Reply{MsgType: message.MsgTypeText,MsgData: text}
	})
	//server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
	//	//TODO
	//	//回复消息：演示回复用户发送的消息
	//	text := message.NewText(msg.Content)
	//	return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	//})
	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		//fmt.Println(err,"-------------------------------------------")
		return
	}
	//println("1111")
	//发送回复的消息
	server.Send()
}

func CheckIsSucess(c *gin.Context) {
	str := c.Query("strData")
	result, err := global.Rdb.Get(str).Result()
	fmt.Println(result, err)
	if err != nil {
		zap.L().Error("GetAccessToken failed", zap.Error(err))
		app.ResponseError(c, app.CodeWxTickerFail)
		return
	}
	user := global.Wx.GetUser()
	if result == "" {
		resultData := bo.MarchLoginData{
			OpenId:"",
			Token:  "",
			Status: 0,
		}
		app.ResponseSuccess(c,resultData)
		return
	}
	info, err := user.GetUserInfo(result)
	//info, err := oauth.GetUserInfo(accessToken, result)
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
	login.LoginService(oauthInfo)
	data, err := login.LoginService(oauthInfo)
	if err != nil {
		zap.L().Error("CreatDaoWeixinUserInfo failed", zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}
	// 如果是statu 是1则请求补全信息接口，if是2则获取到token登录成功
	app.ResponseSuccess(c,data)
}