package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/dto"
	"project/app/march_voiced/service"
	"project/common/database/wx"
	"project/common/global"
	"project/utils/app"
	"strconv"
)

//登录验证获取token
func LoginHandler(c *gin.Context) {
	wx.Init()
	code := c.Query("code")
	status := c.Query("status")
	code ="0414i60w3y7JRV2Qut0w3ZHOUa44i60H"
	fmt.Println(status,"dddd")
	oauth := global.Wx.GetOauth()
	token, err := oauth.GetUserAccessToken(code)
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
	data, err := login.LoginService(info)
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
	data, err:= login.CreatSysUserService(*p)
	if err != nil {
		zap.L().Error("CreatDaoWeixinUserInfo failed", zap.Error(err))
		app.ResponseError(c, app.CodeInsertOperationFail)
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

//测试接口
func Aaa(c *gin.Context) {
	user := new(models.Model)
	login := user.Login()
	c.JSON(http.StatusOK,login)
}
