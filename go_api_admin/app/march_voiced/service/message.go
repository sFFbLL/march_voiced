package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/dto"
)

type Message struct {
}

func (e *Message) MessageUnread(p *dto.MessageUnreadDto, userId int) (count int64, err error) {
	switch p.Type {
	case 1:
		message := new(models.Message)
		count, err = message.UnReadCount()
		if err != nil {
			return
		}
	case 2:
		followMessage := new(models.FollowMessage)
		count, err = followMessage.UnReadCount()
		if err != nil {
			return
		}
	case 3:
		sysMessage := new(models.SysMessage)
		count, err = sysMessage.UnReadCount()
		if err != nil {
			return
		}
	default:
		message := new(models.Message)
		followMessage := new(models.FollowMessage)
		sysMessage := new(models.SysMessage)
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
