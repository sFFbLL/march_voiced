package models

import (
	"errors"
	"fmt"

	"project/app/admin/models"
	"project/app/march_voiced/models/dto"
	"project/common/global"
	"project/utils"

	"github.com/silenceper/wechat/v2/officialaccount/oauth"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//type WxUser struct {
//	Sex           uint8  `json:"sex" gorm:"not null"`                              //性别 0位置 1男 2女
//	IsDeleted     uint8  `json:"is_deleted" gorm:"not null;default:0"`             //0未删除 1已删除
//	IsSubscribe   uint8  `json:"is_subscribe" gorm:"not null"`                     //0未关注 1已关注
//	ID            uint   `json:"id" gorm:"primary_key"`                            //ID
//	SysUserId     uint   `json:"sys_user_id" gorm:"not null;index"`                //后台用户id
//	SubscribeTime uint   `json:"subscribe_time" gorm:"not null"`                   //关注时间
//	CreateTime    uint64 `json:"create_time" gorm:"autoCreateTime:milli;not null"` //创建时间
//	Openid        string `json:"openid" gorm:"not null"`                           //openid
//	WxNickName    string `json:"wx_nick_name" gorm:"not null"`                     //微信昵称
//	Avatar        string `json:"avatar" gorm:"not null"`                           //微信头像
//	Province      string `json:"province" gorm:"not null"`                         //省份
//	City          string `json:"city" gorm:"not null"`                             //市
//	Country       string `json:"country" gorm:"not null"`                          //国家
//	Unionid       string `json:"unionid" gorm:"not null"`                          //微信平台unionid
//}
type WxUser struct {
	Wx_nick_name   string `json:"wxNickName"`
	Openid         string `json:"openid"`
	Sex            int32  `json:"sex" `
	Sys_user_id    int    `json:"sys_user_id"`
	Province       string `json:"province"`
	City           string `json:"city"`
	Country        string `json:"country"`
	Avatar         string `json:"avatar"`
	Unionid        string `json:"unionid"`
	Id             int    `json:"id"`
	Is_subscribe   int    `json:"is_subscribe"`
	Subscribe_time int    `json:"subscribe_time"`
	models.BaseModel
}

type Model struct {
}

type SysUser struct {
	models.BaseModel
	Username     string `json:"username"`
	Password     string `json:"password"`
	DeptId       int    `json:"dept_id"`        //部门id
	PostId       int    `json:"post_id"`        //
	RoleId       int    `json:"role_id"`        //
	NickName     string `json:"nick_name"`      //
	Phone        string `json:"phone"`          //
	Email        string `json:"email"`          //
	AvatarPath   string `json:"avatar_path"`    //头像路径
	Avatar       string `json:"avatar"`         //
	Sex          string `json:"sex"`            //
	Status       string `json:"status"`         //
	Remark       string `json:"remark"`         //
	Salt         string `json:"salt"`           //
	Gender       []byte `json:"gender"`         //性别（0为男默认，1为女）
	IsAdmin      []byte `json:"is_admin"`       //是否为admin账号
	Enabled      []byte `json:"enabled"`        //状态：1启用（默认）、0禁用
	PwdResetTime int64  `json:"pwd_reset_time"` //修改密码的时间
	CreateBy     int    `json:"create_by"`      //
	UpdateBy     int    `json:"update_by"`      //
	Signature    string `json:"signature"`      //个性签名
}

func (a *Model) TableName() string {
	return `wx_user`
}

func (a *Model) LoginDao(info oauth.UserInfo) (int, SysUser, error) {
	var aa WxUser
	var sysuser SysUser
	err := global.Eloquent.First(&aa, "openid=?", info.OpenID).Error

	// 如果用户不存在，创建新用户
	if err == gorm.ErrRecordNotFound {
		newUser := WxUser{
			Wx_nick_name:   info.Nickname,
			Openid:         info.OpenID,
			Sys_user_id:    0,
			Sex:            info.Sex,
			Province:       info.Province,
			City:           info.City,
			Country:        info.Country,
			Avatar:         info.HeadImgURL,
			Unionid:        info.Unionid,
			Is_subscribe:   0,
			Subscribe_time: 0,
		}
		err = global.Eloquent.Create(&newUser).Error
		if err != nil {
			return 0, sysuser, err
		}
		return 1, sysuser, err
	}

	// 查询数据库出错直接返回错误
	if err != nil {
		return 0, sysuser, err
	}

	// 判断用户是否注册
	if aa.Sys_user_id <= 0 {
		return 1, sysuser, nil
	} else {
		err1 := global.Eloquent.First(&sysuser, "id=?", aa.Sys_user_id).Error
		if err1 == gorm.ErrRecordNotFound {
			return 1, sysuser, nil
		}
		if err1 != nil {
			return 0, sysuser, err
		}
	}
	return 2, sysuser, nil
}

func (a *Model) CreatSysUserService(p dto.InsertUserDto) (SysUser, error) {
	var wxUser WxUser
	var isUser SysUser
	var sysUser SysUser
	err := global.Eloquent.First(&wxUser, "openid=?", p.OpenId).Error
	if err != nil {
		return sysUser, err
	}
	err = global.Eloquent.First(&isUser, "username=?", p.UserName).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return sysUser, err
	}
	isUser.Password, err = utils.RsaPriDecode(isUser.Password)
	if err != nil {
		zap.L().Error("ras decode fail", zap.Error(err))
		return SysUser{}, err
	}
	if isUser.Username == p.UserName {
		//判断密码是否相同
		if isUser.Password == p.Password {
			err = global.Eloquent.Model(&wxUser).Where("openid=?", p.OpenId).Update("sys_user_id", isUser.ID).Error
			if err != nil {
				return SysUser{}, err
			}
			return isUser, nil
		} else {
			return SysUser{}, errors.New("该用户名已存在")
		}

	}
	p.Password, err = utils.RsaPubEncode(p.Password)
	if err != nil {
		zap.L().Error("ras encode fail", zap.Error(err))
		return SysUser{}, err
	}
	sysUser = SysUser{
		Username:     p.UserName,
		Password:     p.Password,
		DeptId:       0,
		PostId:       0,
		RoleId:       2,
		NickName:     p.UserName,
		Phone:        "",
		Email:        "",
		AvatarPath:   wxUser.Avatar,
		Avatar:       wxUser.Avatar,
		Sex:          "",
		Status:       "",
		Remark:       "",
		Salt:         "",
		Gender:       []byte{0},
		IsAdmin:      []byte{0},
		Enabled:      []byte{1},
		PwdResetTime: 0,
		CreateBy:     0,
		UpdateBy:     0,
	}
	err = global.Eloquent.Create(&sysUser).Error
	if err != nil {
		return sysUser, err
	}
	err = global.Eloquent.Model(&wxUser).Where("openid=?", p.OpenId).Update("sys_user_id", sysUser.ID).Error
	if err != nil {
		return SysUser{}, err
	}
	return sysUser, nil
}

func (a *Model) SearchUsername(username string) error {
	var user SysUser
	err := global.Eloquent.First(&user, "username=?", username).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else {
		return errors.New("该用户已经存在")
	}
}

func (a *Model) SearchUserInfo(id int) (SysUser, error) {
	var userinfo SysUser
	err := global.Eloquent.First(&userinfo, "id=?", id).Error
	if err != nil {
		return userinfo, err
	}
	return userinfo, nil
}

func (a *Model) ModInformation(userinfo dto.ModInformationDto, id string) error {
	var sex string
	if userinfo.Sex == 1 {
		sex = "男"
	} else {
		sex = "女"
	}
	updateuser := SysUser{
		Username:   userinfo.NickName,
		NickName:   userinfo.NickName,
		AvatarPath: userinfo.AvatarPath,
		Sex:        sex,
		Signature:  userinfo.Signature,
	}
	err := global.Eloquent.Model(SysUser{}).Where("id=?", id).Updates(&updateuser).Error
	return err

}

func (a *Model) Login() WxUser {
	var aa WxUser
	sysUser := SysUser{
		Username:     "111",
		Password:     "123456",
		DeptId:       0,
		PostId:       0,
		RoleId:       0,
		NickName:     "info.Nickname",
		Phone:        "",
		Email:        "",
		AvatarPath:   "info.HeadImgURL",
		Avatar:       "info.HeadImgURL",
		Sex:          "string(info.Sex)",
		Status:       "",
		Remark:       "",
		Salt:         "",
		Gender:       []byte("1"),
		IsAdmin:      []byte{0},
		Enabled:      []byte{1},
		PwdResetTime: 0,
		CreateBy:     0,
		UpdateBy:     0,
	}
	err := global.Eloquent.Create(&sysUser).Error
	if err != nil {
		fmt.Println("创建系统用户失败")
		return aa
	}
	//now := time.Now().Unix()
	return aa
}
