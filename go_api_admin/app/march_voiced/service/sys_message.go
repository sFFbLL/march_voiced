package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/utils"
)

type SysMessage struct {
}

func (e *SysMessage) GetSysMessage(p *dto.Paginator, userId int) (getSysMessage *bo.GetSysMessage, err error) {
	sysMessage := new(models.SysMessage)
	count, message, err := sysMessage.GetSysMessage(p, userId)
	if err != nil {
		return
	}
	getSysMessageData := new([]bo.GetSysMessageData)
	for _, i := range *message {
		*getSysMessageData = append(*getSysMessageData, bo.GetSysMessageData{
			Type:       i.Type,
			Status:     i.Status,
			CreateTime: i.CreateTime,
			Title:      i.Title,
		})
	}
	getSysMessage = new(bo.GetSysMessage)
	getSysMessage.Total = count
	getSysMessage.Records = getSysMessageData
	getSysMessage.Pages = utils.PagesCount(int(count), int(p.Size))
	getSysMessage.Size = p.Size
	getSysMessage.Current = p.Current
	return
}
