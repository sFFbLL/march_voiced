package models

import (
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"
)

type Message struct {
	BaseModel
	CreateBy  uint   `json:"create_by" gorm:"not null;index"`    //创建人
	UpdateBy  uint   `json:"update_by" gorm:"not null"`          //更新人
	ArticleId uint   `json:"article_id" gorm:"not null"`         //文章/三月圈id
	FollowId  uint   `json:"follow_id" gorm:"not null;index"`    //作者id
	Type      uint8  `json:"type" gorm:"not null"`               //消息类型 0文章 1三月圈
	Status    uint8  `json:"status" gorm:"not null"`             //消息状态 文章0发布 1评论 2文章：0发布、1评论、2收藏、3转载、4赞 三月圈：0发布、1评论、2，3，4三种赞
	IsRead    uint8  `json:"is_read" gorm:"not null;default:0"`  //0未读 1已读
	Comment   string `json:"comment" gorm:"not null;default:''"` //评论内容
	Title     string `json:"title" gorm:"not null;default:''"`   //文章题目
	Image     string `json:"image" gorm:"not null;default:''"`   //首图url
	Content   string `json:"content" gorm:"not null;default:''"` //内容
}

func (e *Message) TableName() string {
	return `message`
}

// GetMessageMe 我的消息
func (e *Message) GetMessageMe(getMessage *bo.GetMessage, p *dto.Paginator, userId int) (err error) {
	err = global.Eloquent.Table(e.TableName()).
		Select("sys_user.avatar_path, sys_user.nick_name, message.type , message.status, message.article_id, message.create_by, message.create_time, message.title, message.comment, message.content, message.image").
		Joins("left join sys_user on sys_user.id = message.create_by").
		Where("message.follow_id in (?) and sys_user.is_deleted=0", userId).Count(&getMessage.Total).
		Order("message.create_time desc").Limit(int(p.Size)).Offset(int(p.Current - 1*p.Size)).
		Find(getMessage.Records).Error
	return
}

func (e *Message) GetMessage(getMessage *bo.GetMessage, p *dto.Paginator, userId int) (err error) {
	var followList int
	err = global.Eloquent.Table("follow").Select("follow_id").
		Where("create_by=? and is_deleted=0", userId).Find(&followList).Error
	if err != nil {
		return
	}

	err = global.Eloquent.Table(e.TableName()).
		Select("sys_user.avatar_path, sys_user.nick_name, message.type , message.status, message.article_id, message.create_by, message.create_time, message.title, message.comment, message.content, message.image").
		Joins("left join sys_user on sys_user.id = message.create_by").
		Where("message.create_by in (?) and sys_user.is_deleted=0", followList).Count(&getMessage.Total).
		Order("message.create_time desc").Limit(int(p.Size)).Offset(int(p.Current - 1*p.Size)).
		Find(getMessage.Records).Error
	return
}

// UnRead 未读消息已读业务方法持久方法
func (e *Message) UnRead() error {
	return global.Eloquent.Table(e.TableName()).
		Where("follow_id=? and is_read=? and is_deleted=0", e.FollowId, 0).
		Updates(&Message{IsRead: 1}).Error
}

// UnReadCount 查询是否有未读消息持久方法
func (e *Message) UnReadCount() (count int64, err error) {
	err = global.Eloquent.Table(e.TableName()).
		Where("follow_id=? and is_read=? and is_deleted=0", e.FollowId, 0).
		Count(&count).Error
	return
}

// Add 新增动态消息 t=0文章 t=1三月圈
func (e *Message) Add(articleId uint, userId uint) (err error) {
	if e.Type == 0 {
		article := new(Article)
		err = global.Eloquent.Table("article").Where("id=? and is_deleted=0", articleId).Error
		if err != nil {
			return
		}
		e.ArticleId = uint(article.ID)
		e.FollowId = article.CreateBy
		e.Title = article.Title
		e.Image = article.Image
		e.Content = article.Content
		e.CreateBy = userId
		e.UpdateBy = userId
	} else if e.Type == 1 {
		marchsoft := new(MarchSoft)
		err = global.Eloquent.Table("marchsoft").Where("id=? and is_deleted=0", articleId).Error
		if err != nil {
			return
		}
		e.ArticleId = uint(marchsoft.ID)
		e.FollowId = marchsoft.CreateBy
		e.Image = marchsoft.Image
		e.Content = marchsoft.Content
		e.CreateBy = userId
		e.UpdateBy = userId
	} else {
		return
	}
	return global.Eloquent.Table(e.TableName()).Create(e).Error
}
