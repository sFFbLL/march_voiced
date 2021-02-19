package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/dto"
)

type Message struct {
}

// MessageUnread 查询是否有未读消息业务方法
func (e *Message) MessageUnread(p *dto.MessageUnreadDto, userId int) (count int64, err error) {
	switch p.Type {
	case 1:
		message := new(models.Message)
		message.FollowId = uint(userId)
		count, err = message.UnReadCount()
		if err != nil {
			return
		}
	case 2:
		followMessage := new(models.FollowMessage)
		followMessage.FollowId = uint(userId)
		count, err = followMessage.UnReadCount()
		if err != nil {
			return
		}
	case 3:
		sysMessage := new(models.SysMessage)
		sysMessage.FollowId = uint(userId)
		count, err = sysMessage.UnReadCount()
		if err != nil {
			return
		}
	default:
		message := new(models.Message)
		followMessage := new(models.FollowMessage)
		sysMessage := new(models.SysMessage)
		message.FollowId = uint(userId)
		followMessage.FollowId = uint(userId)
		sysMessage.FollowId = uint(userId)
		var count1, count2, count3 int64
		count1, err = message.UnReadCount()
		if err != nil {
			return
		}
		count2, err = followMessage.UnReadCount()
		if err != nil {
			return
		}
		count3, err = sysMessage.UnReadCount()
		if err != nil {
			return
		}
		count = count1 + count2 + count3
	}
	return
}

// UnReadMessageRead 未读消息已读业务方法
func (e *Message) UnReadMessageRead(p *dto.UnReadMessageRead, userId int) (err error) {
	switch p.Type {
	case 1:
		message := new(models.Message)
		message.FollowId = uint(userId)
		err = message.UnRead()
	case 2:
		followMessage := new(models.FollowMessage)
		followMessage.FollowId = uint(userId)
		err = followMessage.UnRead()
	case 3:
		sysMessage := new(models.SysMessage)
		sysMessage.FollowId = uint(userId)
		err = sysMessage.UnRead()
	default:
		message := new(models.Message)
		followMessage := new(models.FollowMessage)
		sysMessage := new(models.SysMessage)
		message.FollowId = uint(userId)
		followMessage.FollowId = uint(userId)
		sysMessage.FollowId = uint(userId)
		err = message.UnRead()
		if err != nil {
			return
		}
		err = followMessage.UnRead()
		if err != nil {
			return
		}
		err = sysMessage.UnRead()
	}
	return
}
