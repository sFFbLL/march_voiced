package models

import "project/common/global"

type FollowMessage struct {
	BaseModel
	CreateBy uint  `json:"create_by" gorm:"not null;index"`           //创建人
	UpdateBy uint  `json:"update_by" gorm:"not null"`                 //更新人
	FollowId uint  `json:"follow_id" gorm:"not null;default:0;index"` //作者id
	IsRead   uint8 `json:"is_read" gorm:"not null;default:0"`         //0未读 1已读
}

func (e *FollowMessage) TableName() string {
	return `follow_message`
}

func (e *FollowMessage) UnRead() error {
	return global.Eloquent.Table(e.TableName()).
		Where("follow_id=? and is_read=? and is_deleted=0", e.FollowId, 0).
		Updates(&FollowMessage{IsRead: 1}).Error
}

func (e *FollowMessage) UnReadCount() (count int64, err error) {
	err = global.Eloquent.Table(e.TableName()).
		Where("follow_id=? and is_read=? and is_deleted=0", e.FollowId, 0).Count(&count).Error
	return
}
