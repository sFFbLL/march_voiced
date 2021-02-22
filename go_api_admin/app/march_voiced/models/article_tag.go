package models

import (
	"project/app/march_voiced/models/bo"
	"project/common/global"
)

type ArticleTag struct {
	Tag      string `json:"kind" gorm:""`
	Icon     uint   `json:"icon" gorm:""`
	CreateBy uint   `json:"creat_by" gorm:""`
	UpdateBy uint   `json:"update_by" gorm:""`
	BaseModel
}

func (a *ArticleTag) TableName() string {
	return `article_tag`
}

func (a *ArticleTag) GetTagList() (tagList *[]bo.ArticleTagList, err error) {
	tagList = new([]bo.ArticleTagList)

	err = global.Eloquent.Table(a.TableName()).
		Where("is_deleted = 0").Find(tagList).Error
	return
}

func (a *ArticleTag) GetTagById() (tag string, err error) {
	err = global.Eloquent.Table(a.TableName()).Select("tag").
		Where("id = ? AND is_deleted = 0", a.ID).Find(&tag).Error
	return
}
