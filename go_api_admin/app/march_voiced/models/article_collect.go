package models

import (
	"gorm.io/gorm"
	"project/app/march_voiced/models/dto"
	"project/common/global"
)

type ArticleCollect struct {
	ArticleId uint `json:"article_id" gorm:""`
	CreateBy   uint `json:"creat_by" gorm:""`
	UpdateBy  uint `json:"update_by" gorm:""`
	BaseModel
}

func (a *ArticleCollect) TableName() string {
	return `article_collect`
}

func (a *ArticleCollect) AddArticleCollect(p *dto.CollectArticleDto) (err error) {
	table := global.Eloquent.Table(a.TableName())
	var count int64
	err = table.Where("create_by=? and article=? and is_deleted=?", a.CreateBy, a.ArticleId, 0).Count(&count).Error
	if err != nil || count>0 {
		return gorm.ErrInvalidData
	}
	return table.Create(a).Error
}
