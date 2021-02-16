package models

import "project/app/admin/models"

type ArticleTag struct {
	Tag      string `json:"kind" gorm:""`
	Icon     uint   `json:"icon" gorm:""`
	CreatBy  uint   `json:"creat_by" gorm:""`
	UpdateBy uint   `json:"update_by" gorm:""`
	models.BaseModel
}

func (a *ArticleTag) TableName() string {
	return `article_tag`
}
