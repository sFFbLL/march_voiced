package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/dto"
)

type ArticleFavour struct {
}

func (e *ArticleFavour) AddArticleFavour(p *dto.ArticleFavourDto, userId int) (err error) {
	articleFavour := new(models.ArticleFavour)
	articleFavour.ArticleId = p.Id
	articleFavour.UpdateBy = uint(userId)
	articleFavour.CreateBy = uint(userId)
	err = articleFavour.AddArticleFavour(p)
	go articleFavour.AddArticleFavourMessage(uint(userId))
	return
}
