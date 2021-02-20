package models

import (
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"
)

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

// Add 新增关注消息
func (e *FollowMessage) Add() {
	var count int64
	err := global.Eloquent.Table("follow").
		Where("follow_id=? and create_by=? and is_deleted=1", e.FollowId, e.CreateBy).Count(&count).Error
	if err != nil || count > 0{
		return
	}
	_ = global.Eloquent.Table(e.TableName()).Create(e).Error
}

// UnRead 未读消息已读业务方法持久方法
func (e *FollowMessage) UnRead() error {
	return global.Eloquent.Table(e.TableName()).
		Where("follow_id=? and is_read=? and is_deleted=0", e.FollowId, 0).
		Updates(&FollowMessage{IsRead: 1}).Error
}

// UnReadCount 查询是否有未读消息持久方法
func (e *FollowMessage) UnReadCount() (count int64, err error) {
	err = global.Eloquent.Table(e.TableName()).
		Where("follow_id=? and is_read=? and is_deleted=0", e.FollowId, 0).Count(&count).Error
	return
}

func (e *FollowMessage) GetFollowMessage(getFollowMessage *bo.GetFollowMessage, p *dto.Paginator, userId int) (err error) {
	err = global.Eloquent.Table(e.TableName()).
		Select("sys_user.avatar_path, sys_user.nick_name, sys_user.id, follow_message.create_time").
		Joins("left join sys_user on sys_user.id = follow_message.create_by").
		Where("follow_message.follow_id=? and sys_user.is_deleted=0", userId, userId).Count(&getFollowMessage.Total).
		Order("follow_message.create_time desc").
		Limit(int(p.Size)).Offset(int(p.Current - 1*p.Size)).Find(getFollowMessage.Records).Error
	return
}

func (e *FollowMessage) GetUserIsFollow(followList *[]int64, userId int) (err error) {
	err = global.Eloquent.Table("follow").
		Select("follow_id").Where("create_by=?", userId).Find(followList).Error
	return
}
