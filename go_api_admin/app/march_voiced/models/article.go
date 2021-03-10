package models

import (
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"
	"project/utils"

	"go.uber.org/zap"
)

type Article struct {
	Title            string `json:"title" gorm:"size:128;"`
	Content          string `json:"content" gorm:"type(text)"`
	Describe         string `json:"describe" gorm:"size:20"`
	Image            string `json:"image" gorm:"size:128;"`
	Kind             uint8  `json:"kind" gorm:"size:4;"`
	Status           *uint8 `json:"status" gorm:"size:2;"`
	IsRecommend      *uint8 `json:"is_recommend" gorm:"size:1;DEFAULT:0;"`
	Type             *uint  `json:"type" gorm:""`
	Tag              uint   `json:"tag" gorm:""`
	WordCount        uint   `json:"word_count"`
	CreateBy         uint   `json:"create_by" gorm:""`
	UpdateBy         uint   `json:"update_by" gorm:""`
	StatusUpdateTime int64  `json:"status_update_time" gorm:"autoCreateTime:milli"`
	BaseModel
}

func (a *Article) TableName() string {
	return `article`
}

func (a *Article) ArticleSearchContent(data *bo.ArticleCollectByUserId, p *dto.ArticleSearchPaginator, userId int) error {
	return global.Eloquent.Table(a.TableName()).
		Select("article.id, article.title, article.describe, article.image, article.word_count, article.type, article.create_by, article.create_time, sys_user.nick_name").
		Joins("left join sys_user on article.create_by = sys_user.id").
		Where("sys_user.is_deleted=0 and article.is_deleted=0 and article.status=1 and article.content like ?", "%"+p.SearchWord+"%").Count(&data.Total).
		Order("article.create_time desc").Limit(int(p.Size)).Offset(int((p.Current - 1) * p.Size)).
		Find(data.Records).Error
}

func (a *Article) ArticleCollectByUserId(data *bo.ArticleCollectByUserId, p *dto.Paginator, userId int) error {
	return global.Eloquent.Table("article_collect").
		Select("article.id, article.title, article.describe, article.image, article.word_count, article.type, article.create_by, article.create_time, sys_user.nick_name").
		Joins("left join article on article_collect.article_id = article.id").
		Joins("left join sys_user on article.create_by = sys_user.id").
		Where("sys_user.is_deleted=0 and article.is_deleted=0 and article_collect.is_deleted=0 and article.status=1").Count(&data.Total).
		Order("article.create_time desc").Limit(int(p.Size)).Offset(int((p.Current - 1) * p.Size)).
		Find(data.Records).Error
}

// ArticleRecommend
func (a *Article) ArticleRecommend(articleId int) error {
	return global.Eloquent.Table(a.TableName()).Where("id=? and status=1 and is_deleted=0", articleId).First(a).Error
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

// GetApplyArticle （后台）文章审核列表页
func (a *Article) GetApplyArticle(applyArticleList *bo.ApplyArticleList, p *dto.ApplyArticlePaginator, userId int) (err error) {
	nickname := "%" + p.Nickname + "%"
	title := "%" + p.Title + "%"
	table := global.Eloquent.Table(a.TableName()).
		Select("article.id, article.title, article.is_recommend, article.status, article.status_update_time, article_tag.tag, sys_user.nick_name").
		Joins("left join sys_user on sys_user.id = article.create_by").
		Joins("left join article_tag on article_tag.id = article.tag").
		Where("sys_user.is_deleted=0 and article.is_deleted=0 and article.status!=0").
		Where("sys_user.nick_name like ? and article.title like ?", nickname, title)
	if p.Status != 0 && p.Status < 3 {
		table = table.Where("article.status = ?", p.Status)
	}
	if p.Tag != 0 {
		table = table.Where("article.tag = ?", p.Tag)
	}
	if p.EndTime != 0 && p.StartTime != 0 {
		table = table.Where("article.status_update_time > ? AND article.status_update_time < ?", p.StartTime, p.EndTime)
	}
	err = table.Count(&applyArticleList.Total).
		Order("article.status_update_time desc").Limit(int(p.Size)).Offset(int((p.Current - 1) * p.Size)).
		Find(applyArticleList.Records).Error
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

func (a *Article) DeleteArticle() (err error) {
	err = global.Eloquent.Table(a.TableName()).Where("id = ? AND is_deleted = 0", a.ID).Updates(a).Error
	if err != nil {
		zap.L().Error("DeleteArticle Dao Update Failed", zap.Error(err))
	}
	return
}

func (a *Article) ArticleDetail() (articleMsg *bo.ArticleDetail, err error) {
	articleMsg = new(bo.ArticleDetail)
	// 获取文章信息
	err = global.Eloquent.Table(a.TableName()).
		Select("article.id, article.title, article.content, article.image, article.word_count,article.status, article.kind, article.type, article.create_time, article.create_by, article.update_by, article.update_time, article_tag.tag, sys_user.nick_name, sys_user.avatar_path").
		Joins("JOIN sys_user ON article.create_by = sys_user.id").
		Joins("JOIN article_tag ON article.tag = article_tag.id").
		Where("article.id = ? AND article.is_deleted = 0", a.ID).
		First(articleMsg).Error
	return
}

func (a *Article) ArticleList(paging dto.Paging, IsRecommend int) (articleArray *[]bo.Article, err error) {
	articleArray = new([]bo.Article)
	err = global.Eloquent.Table(a.TableName()).
		Select("article.id, article.title, article.describe, article.image, article.status, article.type, article.create_time, article.create_by, article.update_by, article.update_time, sys_user.nick_name, sys_user.avatar_path").
		Joins("JOIN sys_user ON article.create_by = sys_user.id").
		Where("article.is_recommend = ? AND article.is_deleted = 0 AND article.status = 1", IsRecommend).
		Limit(int(paging.Size)).Offset(int((paging.Current - 1) * paging.Size)).Find(articleArray).Error
	return
}

func (a *Article) SelectArticleListByUserId(paging dto.SelectArticleByUser) (articleArray *[]bo.ArticleMsg, err error) {
	articleArray = new([]bo.ArticleMsg)
	if paging.Kind == 1 {
		err = global.Eloquent.Table(a.TableName()).
			Where("create_by = ? AND status = ? AND is_deleted = 0", paging.ID, paging.Kind).
			Limit(int(paging.Size)).Offset(int((paging.Current - 1) * paging.Size)).Find(articleArray).Error
		return
	}
	err = global.Eloquent.Table(a.TableName()).
		Where("create_by = ? AND is_deleted = 0", paging.ID).
		Limit(int(paging.Size)).Offset(int((paging.Current - 1) * paging.Size)).Find(articleArray).Error
	return
}
