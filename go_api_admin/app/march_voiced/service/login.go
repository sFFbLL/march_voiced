package service

import (
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/pkg/jwt"
)

type Login struct {
}
func(login *Login) LoginService(info oauth.UserInfo) (data bo.MarchLoginData,err error) {
	user := new(models.Model)
	status, sysUser, err := user.LoginDao(info)
	if err != nil {
		return
	}
	var token string
	if status == 2{
		token, err =jwt.GenToken(sysUser.ID, sysUser.Username)
		if err != nil {
			return data, err
		}
		token = "Bearer"+" "+token
	}
	data = bo.MarchLoginData{
		OpenId:info.OpenID,
		Token:  token,
		Status: status,
	}
	return data, err
}

func(login *Login) CreatSysUserService(p dto.InsertUserDto) (data bo.LoginInfoData, err error){
	user := new(models.Model)
	sysUser, err := user.CreatSysUserService(p)
	if sysUser.ID > 0{
		data.Token, err =jwt.GenToken(sysUser.ID, sysUser.Username)
		if err != nil {
			return
		}
		data.Token = "Bearer "+ data.Token
	}
	return
}

func (login *Login) SearchUsername(username string) error {
	user := new(models.Model)
	err := user.SearchUsername(username)
	return err
}

func (login *Login) SearchUserInfo(id int)(bo.UserInfo, error){
	user := new(models.Model)
	var userinfo bo.UserInfo
	var sex int
	info, err := user.SearchUserInfo(id)
	if info.Sex =="男"{
		sex = 0
	}else if info.Sex =="女"{
		sex =1
	}
	userinfo = bo.UserInfo{
		NickName:   info.NickName,
		Sex:        sex,
		Signature:  info.Signature,
		AvatarPath: info.AvatarPath,
	}
	//strconv.Atoi(info.Gender),
	return userinfo, err
}

func (login *Login)ModInformation(userinfo dto.ModInformationDto,id string) error {
	user := new(models.Model)
	err := user.ModInformation(userinfo, id)
	return err
}