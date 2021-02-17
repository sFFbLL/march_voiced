package models

type FollowMessage struct {
	BaseModel
	CreateBy   uint  `json:"create_by" gorm:"not null;index"`              //创建人
	UpdateBy   uint  `json:"update_by" gorm:"not null"`                    //更新人
	AuthUserID uint  `json:"auth_user_id" gorm:"not null;default:0;index"` //作者id
	IsRead     uint8 `json:"is_read" gorm:"not null;default:0"`            //0未读 1已读
}

func (e *FollowMessage) TableName() string {
	return `follow_message`
}
