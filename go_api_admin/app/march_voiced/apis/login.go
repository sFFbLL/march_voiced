package apis

import (
	"bytes"
	"errors"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/dto"
	"project/app/march_voiced/service"
	"project/common/database/wx"
	"project/common/global"
	"project/utils"
	"project/utils/app"
	"strconv"
	"time"
)

//登录验证获取token
func LoginHandler(c *gin.Context) {
	wx.Init()
	type Name struct {
		Code string `json:"code"`
	}
	p := new(Name)
	c.ShouldBindJSON(p)
	oauth := global.Wx.GetOauth()
	token, err := oauth.GetUserAccessToken(p.Code)
	if err != nil {
		zap.L().Error("GetAccessToken failed", zap.Error(err))
		app.ResponseError(c, app.CodeWxOuttime)
		return
	}
	info, err := oauth.GetUserInfo(token.AccessToken, token.OpenID)
	fmt.Println(info)
	if err != nil {
		zap.L().Error("GetWeixinUserInfo failed", zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}
	login := new(service.Login)
	data, err := login.LoginService(c,info)
	if err != nil {
		zap.L().Error("CreatDaoWeixinUserInfo failed", zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
		return
	}
	// 如果是statu 是1则请求补全信息接口，if是2则获取到token登录成功
	app.ResponseSuccess(c,data)
}

// 根据oppid创建新的用户
func CreatSysUser(c *gin.Context){
	// 1.获取参数 校验参数
	p := new(dto.InsertUserDto)
	if err := c.ShouldBindJSON(p); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	login := new(service.Login)
	data, err:= login.CreatSysUserService(c,*p)
	if err != nil{
		zap.L().Error("SearchUserName failed", zap.Error(err))
		app.ResponseErrorWithMsg(c, app.CodeOperationFail,err.Error())
		return
	}
	app.ResponseSuccess(c,data)
}

// 查询用户名是否唯一
func SearchUsername(c *gin.Context){
	username := c.Query("username")
	if username ==""{
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	checkUsername := new(service.Login)
	err := checkUsername.SearchUsername(username)
	if err != nil {
		zap.L().Error("SearchUserName failed", zap.Error(err))
		app.ResponseError(c, app.CodeUserNameExist)
		return
	}
	app.ResponseSuccess(c,app.CodeSuccess)
}

// 查询用户信息
func SearchUserInfo(c *gin.Context){
	id := c.Query("id")
	if id ==""{
		get, _ := c.Get("UserId")
		id = fmt.Sprintf("%v",get)
	}
	atoi, err := strconv.Atoi(id)
	if err != nil && atoi <=0{
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	searchUserinfo := new(service.Login)
	data,err := searchUserinfo.SearchUserInfo(atoi)
	if err != nil {
		zap.L().Error("SearchUserInfo failed", zap.Error(err))
		app.ResponseError(c, app.CodeSelectOperationFail)
		return
	}
	app.ResponseSuccess(c,data)
}

func ModInformation(c *gin.Context)  {
	get, _ := c.Get("UserId")
	id := fmt.Sprintf("%v",get)
	p := new(dto.ModInformationDto)
	if err := c.ShouldBindJSON(p); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	user := new(service.Login)
	err := user.ModInformation(*p, id)
	if err != nil {
		zap.L().Error("ModUserName failed", zap.Error(err))
		app.ResponseError(c, app.CodeUpdateOperationFail)
		return
	}
	app.ResponseSuccess(c,app.CodeSuccess)
}

//生成带参数的二维码
func GetTicket(c *gin.Context) {
	var accessToken string
	//wxAccessToken
	token, err := global.Rdb.Get("wxAccessToken").Result()
	if err != nil {
		access, err := utils.GetAccess()
		if err != nil {
			app.ResponseError(c, app.CodeWxTickerFail)
			return
		}
		global.Rdb.Set("wxAccessToken",access,300*time.Second)
		accessToken = access
	}else{
		accessToken = token
	}
	//fmt.Println(accessToken,"-----------------")
	wxTicket, err := WxTicket(accessToken)
	if err != nil {
		app.ResponseError(c, app.CodeWxTickerFail )
	}
	app.ResponseSuccess(c,wxTicket)
}

func WxTicket(token string)(dto.WxTokenMessages,error){
	//token := "42_bxJkHtccnA6EnfqnQjzC8z3SfoxUsZd9gTEw6oKW_wCDPa3bqJ-6O_D-byqt1HXhLuiJIGZhPp1vVJO-tnCDmjR35Jv-ht8vFdPZo6eQtXgt_UurAhni7-EPhJkFvhwPpd0Juuhs2RAQvl0NPZLcAEAVFR"
	//sceneStr:="test"
	post := `{
	"expire_seconds":300,
	"action_name": "QR_LIMIT_STR_SCENE",
		"action_info": {
		"scene": {
			"scene_str": "%s"
			}
		}
	}`
	str := string(utils.Krand(16, utils.KC_RAND_KIND_ALL))
	_, err1 := global.Rdb.Get(str).Result()
	if err1 != nil {
		global.Rdb.Set(str, "", 300*time.Second)
		aa :=fmt.Sprintf(post,str)
		var ticket dto.WxTokenMessages
		var jsonstr = []byte(aa) //转换二进制
		buffer:= bytes.NewBuffer(jsonstr)
		request, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="+token, buffer)
		if err != nil {
			fmt.Printf("http.NewRequest%v", err)
			return ticket,err
		}
		client := http.Client{} //创建客户端
		resp, err := client.Do(request.WithContext(context.TODO())) //发送请求
		if err != nil {
			fmt.Printf("client.Do%v", err)
			return ticket,err
		}

		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ioutil.ReadAll%v", err)
			return ticket,err
		}
		err = json.Unmarshal(respBytes, &ticket)
		if err != nil {
			fmt.Println(err)
			return ticket,err
		}
		ticket.Strdata = str
		fmt.Println(ticket)
		return ticket ,nil
	}else{
		return dto.WxTokenMessages{},errors.New("请重新请求")
	}
}
//测试接口
func Aaa(c *gin.Context) {
	user := new(models.Model)
	login := user.Login()
	c.JSON(http.StatusOK,login)
}
