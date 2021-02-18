package models

import (
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"

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
	CreatBy     uint   `json:"creat_by" gorm:""`
	UpdateBy    uint   `json:"update_by" gorm:""`
	BaseModel
}

func (a *Article) TableName() string {
	return `article`
}

func (a *Article) InsertArticle() (err error) {
	err = global.Eloquent.Table(a.TableName()).Create(a).Error
	return
}

func (a *Article) UpdateArticle(id uint) (err error) {
	update := new(Article)

	err = global.Eloquent.Table(a.TableName()).Where("id = ? AND is_deleted = 0", id).First(update).Error
	if err != nil {
		zap.L().Error("UpdateArticle Dao Find First Failed", zap.Error(err))
		return
	}

	err = global.Eloquent.Table(a.TableName()).Model(update).Updates(a).Error
	if err != nil {
		zap.L().Error("UpdateArticle Dao Update Failed", zap.Error(err))
	}
	return
}

func (a *Article) ArticleDetail(id int, userId int) (userMsg bo.UserMsg, total [3]int64, err error) {
	totalString := [3]string{"article_collect", "article_comment", "article_favour"}

	// 获取文章信息
	err = global.Eloquent.Table(a.TableName()).Where("id = ? AND is_deleted = 0", id).First(a).Error
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
	err = global.Eloquent.Table("sys_user").Select("id, nick_name, avatar_path").Where("id = ? AND is_deleted = 0", a.CreatBy).First(&userMsg).Error
	if err != nil {
		zap.L().Error("ArticleDetail Select articleUser failed", zap.Error(err))
		return
	}

	// 是否关注
	if uint(userId) != a.CreatBy {
		var count int64
		err = global.Eloquent.Table("follow").Where("follow_id = ? AND creat_by = ? AND is_deleted = 0", a.CreatBy, userId).Count(&count).Error
		if err != nil {
			zap.L().Error("ArticleDetail Select IsFollow failed", zap.Error(err))
			return
		}
		if count > 0 {
			userMsg.IsFollow = 1
		}
	}

	// 获取文章total信息
	for i, v := range totalString {
		err = global.Eloquent.Table(v).Where("article_id = ? AND is_deleted = 0", id).Count(&total[i]).Error
		if err != nil {
			zap.L().Error("ArticleDetail Select total "+v+" failed", zap.Error(err))
			return
		}
	}

	return
}

func (a *Article) TopArticleList(paging dto.Paging) (articleArray *[]Article, err error) {
	articleArray = new([]Article)
	err = global.Eloquent.Table(a.TableName()).Where("is_deleted = 0 AND is_recommend = 1 AND type = 0 AND status = 1").
		Limit(paging.Size).Offset((paging.Current - 1) * paging.Size).Find(articleArray).Error
	return
}
