package models

import (
	"errors"

	"project/app/march_voiced/models/dto"
	"project/common/global"
)

type ArticleFavour struct {
	ArticleId uint `json:"article_id" gorm:""`
	CreateBy  uint `json:"creat_by" gorm:""`
	UpdateBy  uint `json:"update_by" gorm:""`
	BaseModel
}

func (a *ArticleFavour) TableName() string {
	return `article_favour`
}

func (a *ArticleFavour) AddArticleFavour(p *dto.ArticleFavourDto) (err error) {
	table := global.Eloquent.Table(a.TableName())
	var count int64
	err = table.Where("create_by=? and article_id=? and is_deleted=?", a.CreateBy, a.ArticleId, 0).Count(&count).Error
	if err != nil {
		return errors.New("查询操作失败")
	}
	if count>0 {
		a.IsDeleted = 1
		return table.Where("create_by=? and article_id=? and is_deleted=?", a.CreateBy, a.ArticleId, 0).Updates(a).Error
	}
	return table.Create(a).Error
}

func (a *ArticleFavour) AddArticleFavourMessage(userId uint) {
	var count int64
	err := global.Eloquent.Table(a.TableName()).Where("create_by=? and article_id=? and is_deleted=1", a.CreateBy, a.ArticleId).Count(&count).Error
	if err != nil || count > 0 {
		return
	}
	AddMessage(0, 4, a.ArticleId, userId, "")
}

func (a *ArticleFavour) ArticleFavourCount() (count int64, err error) {
	err = global.Eloquent.Table(a.TableName()).
		Where("article_id=? and is_deleted=0", a.ArticleId).Count(&count).Error
	return
}

