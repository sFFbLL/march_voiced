package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/dto"
)

type MarchFavour struct {
}

func (e *MarchFavour) AddMarchFavour(p *dto.MarchFavourDto, userId int) (err error) {
	marchFavour := new(models.MarchSoftFavour)
	marchFavour.MarchsoftId = p.Id
	marchFavour.Type = *p.Type
	marchFavour.UpdateBy = uint(userId)
	marchFavour.CreateBy = uint(userId)
	err = marchFavour.AddMarchFavour(p)
	go marchFavour.AddMarchFavourMessage(p.Type, uint(userId))
	return
}