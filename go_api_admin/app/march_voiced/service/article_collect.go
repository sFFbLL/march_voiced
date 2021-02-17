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
	articleCollect.CreateBy = uint(userId)
	articleCollect.UpdateBy = uint(userId)
	err = articleCollect.AddArticleCollect(p)
	return
}
