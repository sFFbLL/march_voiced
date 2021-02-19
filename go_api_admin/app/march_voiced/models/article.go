package models

import (
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"
	"project/utils"

	"go.uber.org/zap"
)

type Article struct {
	Title       string `json:"title" gorm:"size:128;"`
	Content     string `json:"content" gorm:"type(text)"`
	Image       string `json:"image" gorm:"size:128;"`
	Tag         uint8  `json:"tag" gorm:"size:5;"`
	Kind        uint8  `json:"kind" gorm:"size:4;"`
	Status      *uint8 `json:"status" gorm:"size:2;"`
	IsRecommend *uint8 `json:"is_recommend" gorm:"size:1;DEFAULT:0;"`
	Type        uint   `json:"type" gorm:""`
	CreateBy    uint   `json:"create_by" gorm:""`
	UpdateBy    uint   `json:"update_by" gorm:""`
	BaseModel
}

func (a *Article) TableName() string {
	return `article`
}

func (a *Article) ArticleCountByUserId(id int) (count int64, err error) {
	err = global.Eloquent.Table(a.TableName()).Where("create_by = ? AND is_deleted = 0", id).Count(&count).Error
	if err != nil {
		zap.L().Error("ArticleCountByUserId "+utils.IntToString(id)+" Failed", zap.Error(err))
	}
	return
}

func (a *Article) GetArticle() (err error) {
	err = global.Eloquent.Table(a.TableName()).Where("id = ? AND is_deleted = 0", a.ID).First(a).Error
	if err != nil {
		zap.L().Error("GetArticle Dao Find First Failed", zap.Error(err))
	}
	return
}

func (a *Article) InsertArticle() (err error) {
	err = global.Eloquent.Table(a.TableName()).Create(a).Error
	return
}

func (a *Article) UpdateArticle() (err error) {
	err = global.Eloquent.Table(a.TableName()).Where("id = ? AND is_deleted = 0", a.ID).Updates(a).Error
	if err != nil {
		zap.L().Error("UpdateArticle Dao Update Failed", zap.Error(err))
	}
	return
}

func (a *Article) ArticleDetail() (userMsg bo.UserMsg, err error) {

	// 获取文章信息
	err = global.Eloquent.Table(a.TableName()).Where("id = ? AND is_deleted = 0", a.ID).First(a).Error
	if err != nil {
		zap.L().Error("ArticleDetail Select article failed", zap.Error(err))
		return
	}

	// 是转发的话
	if a.Type != 0 {
		err = global.Eloquent.Table(a.TableName()).Where("id = ? AND is_deleted = 0", a.Type).First(a).Error
		if err != nil {
			zap.L().Error("ArticleDetail Select type article failed", zap.Error(err))
			return
		}
	}

	// 获取用户信息
	err = global.Eloquent.Table("sys_user").Select("sys_user.id AS user_id, nick_name, avatar_path").Where("id = ? AND is_deleted = 0", a.CreateBy).Find(&userMsg).Error
	if err != nil {
		zap.L().Error("ArticleDetail Select articleUser failed", zap.Error(err))
		return
	}

	return
}

func (a *Article) TopArticleList(paging dto.Paging) (articleArray *[]bo.Article, err error) {
	articleArray = new([]bo.Article)
	err = global.Eloquent.Table(a.TableName()).
		Where("article.is_deleted = 0 AND article.is_recommend = 1 AND article.type = 0 AND article.status = 1").
		Select("article.id, article.title, article.content, article.image, article.tag, article.kind, article.type, article.create_time, article.create_by, article.update_by, article.update_time, sys_user.id AS user_id, sys_user.nick_name, sys_user.avatar_path").
		Joins("JOIN sys_user ON article.create_by = sys_user.id").
		Limit(paging.Size).Offset((paging.Current - 1) * paging.Size).Find(articleArray).Error
	return
}
