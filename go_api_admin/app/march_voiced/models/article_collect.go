package models

import "project/app/admin/models"

type ArticleCollect struct {
	ArticleId uint `json:"article_id" gorm:""`
	CreatBy   uint `json:"creat_by" gorm:""`
	UpdateBy  uint `json:"update_by" gorm:""`
	models.BaseModel
}

func (a *ArticleCollect) TableName() string {
	return `article_collect`
}
