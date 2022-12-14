package models

import (
	"project/common/global"
)

// user实体
type User struct {
	BaseModel
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
	IsMarch      int    `json:"is_march"`
	Signature    string `json:"signature"` // 个签
}

type UserIn struct {
	Id         int    `json:"id"`
	AvatarPath string `json:"avatar_path"`
	NickName   string `json:"nick_name"`
	IsFollow   int    `json:"isFollow"`
	Signature  string `json:"signature"`
}

// SysUser用户表名
func (u *User) TableName() string {
	return `sys_user`
}

// SelectUserInfo 查询用户基本信息
func (u *User) SelectUserInfo(id int) (userIn *UserIn, err error) {
	userIn = new(UserIn)
	err = global.Eloquent.Table(u.TableName()).Where("is_deleted=? AND id=?", []byte{0}, id).Find(userIn).Error
	if err != nil {
		return nil, err
	}
	return
}

// GetUserCollect 用戶的文章收藏數
func (u *User) GetUserCollect(id int) (count int64, err error) {
	err = global.Eloquent.Table(`article_collect`).Where("is_deleted=? AND create_by=?", []byte{0}, id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return
}

// 是否在三月圈
func (u *User) IsMarchByUserId() (signal uint8, err error) {
	err = global.Eloquent.Table(u.TableName()).Select("is_march").
		Where("id = ? AND is_deleted = 0 ", u.ID).Find(&signal).Error
	return
}
