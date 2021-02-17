package models

type SysMessage struct {
	BaseModel
	CreateBy   uint  `json:"create_by" gorm:"not null;index"`   //创建人
	UpdateBy   uint  `json:"update_by" gorm:"not null"`         //更新人
	AuthUserID uint   `json:"auth_user_id" gorm:"not null"`      //作者id
	IsRead     uint8 `json:"is_read" gorm:"not null;default:0"` //0未读 1已读
	Type       uint8 `json:"type" gorm:"not null"`              //0文章 1三月圈
	Status     uint8 `json:"status" gorm:"not null;default:0"`  //0待审核 1通过 2驳回
}

func (e *SysMessage) TableName() string {
	return `sys_message`
}
