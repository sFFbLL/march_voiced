package models

import "project/common/global"

type SysMessage struct {
	BaseModel
	CreateBy uint  `json:"create_by" gorm:"not null;index"`           //创建人
	UpdateBy uint  `json:"update_by" gorm:"not null"`                 //更新人
	FollowId uint  `json:"follow_id" gorm:"not null;default:0;index"` //作者id
	IsRead   uint8 `json:"is_read" gorm:"not null;default:0"`         //0未读 1已读
	Type     uint8 `json:"type" gorm:"not null"`                      //0文章 1三月圈
	Status   uint8 `json:"status" gorm:"not null;default:0"`          //0驳回 1待审核 2通过
}

func (e *SysMessage) TableName() string {
	return `sys_message`
}

func (e *SysMessage) UnRead() error {
	return global.Eloquent.Table(e.TableName()).
		Where("follow_id=? and is_read=? and is_deleted=0", e.FollowId, 0).
		Updates(&SysMessage{IsRead: 1}).Error
}

func (e *SysMessage) UnReadCount() (count int64, err error) {
	err = global.Eloquent.Table(e.TableName()).
		Where("follow_id=? and is_read=? and is_deleted=0", e.FollowId, 0).Count(&count).Error
	return
}

func (e *SysMessage) Add() {
	_ = global.Eloquent.Table(e.TableName()).Create(e).Error
}
