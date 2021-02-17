package models

type WxUser struct {
	Sex           uint8  `json:"sex" gorm:"not null"`                              //性别 0位置 1男 2女
	IsDeleted     uint8  `json:"is_deleted" gorm:"not null;default:0"`             //0未删除 1已删除
	IsSubscribe   uint8  `json:"is_subscribe" gorm:"not null"`                     //0未关注 1已关注
	ID            uint   `json:"id" gorm:"primary_key"`                            //ID
	SysUserId     uint   `json:"sys_user_id" gorm:"not null;index"`                //后台用户id
	SubscribeTime uint   `json:"subscribe_time" gorm:"not null"`                   //关注时间
	CreateTime    uint64 `json:"create_time" gorm:"autoCreateTime:milli;not null"` //创建时间
	Openid        string `json:"openid" gorm:"not null"`                           //openid
	WxNickName    string `json:"wx_nick_name" gorm:"not null"`                     //微信昵称
	Avatar        string `json:"avatar" gorm:"not null"`                           //微信头像
	Province      string `json:"province" gorm:"not null"`                         //省份
	City          string `json:"city" gorm:"not null"`                             //市
	Country       string `json:"country" gorm:"not null"`                          //国家
	Unionid       string `json:"unionid" gorm:"not null"`                          //微信平台unionid
}

func (e *WxUser) TableName() string {
	return `wx_user`
}
