package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/utils"
)

type FollowMessage struct {
}

// GetFollowMessage 获取关注消息业务
func (e *FollowMessage) GetFollowMessage(p *dto.Paginator, userId int) (getFollowMessage *bo.GetFollowMessage, err error) {
	followMessage := new(models.FollowMessage)

	getFollowMessage = new(bo.GetFollowMessage)
	followMessageData := new([]bo.GetFollowMessageData)
	getFollowMessage.Records = followMessageData
	err = followMessage.GetFollowMessage(getFollowMessage, p, userId)
	if err != nil {
		return
	}

	// 我关注的人
	followList := new([]int64)
	err = followMessage.GetUserIsFollow(followList, userId)
	if err != nil {
		return
	}
	m := make(map[int64]int64, len(*followList))
	for _, i := range *followList {
		m[i] = 0
	}

	followMessageDataFinal := new([]bo.GetFollowMessageData)
	for _, i := range *followMessageData {
		if _, ok := m[i.Id]; ok {
			i.IsFollow = 1
		}
		*followMessageDataFinal = append(*followMessageDataFinal, i)
	}
	getFollowMessage.Records = followMessageDataFinal
	getFollowMessage.Pages = utils.PagesCount(int(getFollowMessage.Total), int(p.Size))
	getFollowMessage.Current = p.Current
	getFollowMessage.Size = p.Size
	return
}
