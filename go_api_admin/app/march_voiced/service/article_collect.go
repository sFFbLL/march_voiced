package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/dto"
)

type ArticleCollect struct {
}

func (e *ArticleCollect) AddArticleCollect(p *dto.CollectArticleDto, userId int) (err error) {
	articleCollect := new(models.ArticleCollect)
	articleCollect.ArticleId = p.Id
	articleCollect.UpdateBy = uint(userId)
	articleCollect.CreateBy = uint(userId)
	err = articleCollect.AddArticleCollect(p)
	go articleCollect.AddArticleCollectMessage(uint(userId))
	return
}
