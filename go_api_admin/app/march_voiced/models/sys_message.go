package models

import (
	"project/app/march_voiced/models/dto"
	"project/common/global"
)

type SysMessage struct {
	BaseModel
	CreateBy uint   `json:"create_by" gorm:"not null;index"`           //创建人
	UpdateBy uint   `json:"update_by" gorm:"not null"`                 //更新人
	FollowId uint   `json:"follow_id" gorm:"not null;default:0;index"` //作者id
	IsRead   uint8  `json:"is_read" gorm:"not null;default:0"`         //0未读 1已读
	Type     uint8  `json:"type" gorm:"not null"`                      //0文章 1三月圈
	Status   uint8  `json:"status" gorm:"not null;default:0"`          //0驳回 1待审核 2通过
	Title    string `json:"title" gorm:"not null"`
}

func (e *SysMessage) TableName() string {
	return `sys_message`
}

// UnRead 未读消息已读业务方法持久方法
func (e *SysMessage) UnRead() error {
	return global.Eloquent.Table(e.TableName()).
		Where("follow_id=? and is_read=? and is_deleted=0", e.FollowId, 0).
		Updates(&SysMessage{IsRead: 1}).Error
}

// UnReadCount 查询是否有未读消息持久方法
func (e *SysMessage) UnReadCount() (count int64, err error) {
	err = global.Eloquent.Table(e.TableName()).
		Where("follow_id=? and is_read=? and is_deleted=0", e.FollowId, 0).Count(&count).Error
	return
}

// Add 新增系统消息
func (e *SysMessage) Add() {
	_ = global.Eloquent.Table(e.TableName()).Create(e).Error
}

func (e *SysMessage) GetSysMessage(p *dto.Paginator, userId int) (count int64, sysMessage *[]SysMessage, err error) {
	sysMessage = new([]SysMessage)
	err = global.Eloquent.Table(e.TableName()).Where("follow_id", userId).Count(&count).
		Order("create_time desc").
		Limit(int(p.Size)).Offset(int((p.Current - 1) * p.Size)).Find(sysMessage).Error
	return
}
