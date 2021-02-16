package models

import "project/app/admin/models"

type ArticleFavour struct {
	ArticleId uint `json:"article_id" gorm:""`
	CreatBy   uint `json:"creat_by" gorm:""`
	UpdateBy  uint `json:"update_by" gorm:""`
	models.BaseModel
}

func (a *ArticleFavour) TableName() string {
	return `article_favour`
}
