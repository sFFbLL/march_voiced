package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/utils"
	"sync"
)

type Message struct {
}

// GetMessageMe 我的消息
func (e *Message) GetMessageMe(p *dto.Paginator, userId int) (data *bo.GetMessage, err error) {
	message := new(models.Message)
	data = new(bo.GetMessage)
	records := new([]bo.GetMessageData)
	data.Records = records
	err = message.GetMessageMe(data, p, userId)
	data.Pages = utils.PagesCount(int(data.Total), int(p.Size))
	return
}

// GetMessage 获取关注的人动态业务方法
func (e *Message) GetMessage(p *dto.Paginator, userId int) (data *bo.GetMessage, err error) {
	message := new(models.Message)
	data = new(bo.GetMessage)
	records1 := new([]bo.GetMessageData)
	records2 := new([]bo.GetMessageData)

	var goArticle bo.GoArticleMsg
	data.Records = records1
	err = message.GetMessage(data, p, userId)
	data.Pages = utils.PagesCount(int(data.Total), int(p.Size))

	articleMapList := make(map[uint]*bo.GetMessageData, len(*data.Records))
	for _, i := range *data.Records {
		var articleBo = i
		articleMapList[i.Id] = &articleBo
	}
	// 数据拼接
	var wg sync.WaitGroup
	articleCh := make(chan *bo.GoArticleMsg, len(*data.Records))
	for _, v := range *data.Records {
		wg.Add(1)
		goArticle.ArticleId = v.ArticleId
		goArticle.ArticleUserId = v.UserId
		goArticle.UserId = v.UserId
		goArticle.MessageId = v.Id
		go goArticleMsg(&articleCh, &wg, goArticle)
	}
	wg.Wait()
	close(articleCh)

	for i := range articleCh {
		articleMapList[i.MessageId].ArticleTotal = goArticle.ArticleTotal
	}

	for _, i := range *data.Records {
		*records2 = append(*records2, *articleMapList[i.Id])
	}
	data.Records = records2
	return
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
