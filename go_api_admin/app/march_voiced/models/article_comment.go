package models

import "project/app/admin/models"

type ArticleComment struct {
	ArticleId uint `json:"article_id" gorm:""`
	Pid       uint `json:"pid" gorm:""`
	ReplyId   uint `json:"reply_id" gorm:""`
	CreatBy   uint `json:"creat_by" gorm:""`
	UpdateBy  uint `json:"update_by" gorm:""`
	models.BaseModel
}

func (a *ArticleComment) TableName() string {
	return `article_comment`
}
