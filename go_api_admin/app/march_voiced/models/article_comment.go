package models

import (
	"project/common/global"

	"go.uber.org/zap"
)

type ArticleComment struct {
	ArticleId uint   `json:"article_id" gorm:""`
	Pid       uint   `json:"pid" gorm:""`
	ReplyId   uint   `json:"reply_id" gorm:""`
	CreateBy  uint   `json:"creat_by" gorm:""`
	UpdateBy  uint   `json:"update_by" gorm:""`
	Content   string `json:"content" gorm:""`
	BaseModel
}

func (a *ArticleComment) TableName() string {
	return `article_comment`
}

func (a *ArticleComment) ArticleCommentCount(id int) (count int64, err error) {
	err = global.Eloquent.Table(a.TableName()).Where("article_id = ? AND is_deleted = 0", id).Count(&count).Error
	if err != nil {
		zap.L().Error("ArticleCommentCount Select failed", zap.Error(err))
		return
	}
	return
}
