package models

import "project/common/global"

type Message struct {
	BaseModel
	CreateBy   uint   `json:"create_by" gorm:"not null;index"`    //创建人
	UpdateBy   uint   `json:"update_by" gorm:"not null"`          //更新人
	ArticleId  uint   `json:"article_id" gorm:"not null"`         //文章/三月圈id
	AuthUserID uint   `json:"auth_user_id" gorm:"not null;index"` //作者id
	Type       uint8  `json:"type" gorm:"not null"`               //消息类型 0文章 1三月圈
	Status     uint8  `json:"status" gorm:"not null"`             //消息状态 文章0发布 1评论 2文章：0发布、1评论、2收藏、3转载、4赞 三月圈：0发布、1评论、2，3，4三种赞
	IsRead     uint8  `json:"is_read" gorm:"not null;default:0"`  //0未读 1已读
	Comment    string `json:"comment" gorm:"not null;default:''"` //评论内容
	Title      string `json:"title" gorm:"not null;default:''"`   //文章题目
	Image      string `json:"image" gorm:"not null;default:''"`   //首图url
	Content    string `json:"content" gorm:"not null;default:''"` //内容
}

func (e *Message) TableName() string {
	return `message`
}

func (e *Message) UnRead() error {
	return global.Eloquent.Table(e.TableName()).Updates(&Message{IsRead: 1}).Error
}

func (e *Message) UnReadCount() (count int64, err error) {
	err = global.Eloquent.Table(e.TableName()).Where("is_read=?", 0).Count(&count).Error
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
		e.AuthUserID = article.CreatBy
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
		e.AuthUserID = marchsoft.CreateBy
		e.Image = marchsoft.Image
		e.Content = marchsoft.Content
		e.CreateBy = userId
		e.UpdateBy = userId
	} else {
		return
	}
	return global.Eloquent.Table(e.TableName()).Create(e).Error
}
