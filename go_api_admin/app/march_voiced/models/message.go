package models

type Message struct {
	BaseModel
	CreateBy   uint   `json:"create_by" gorm:"not null;index"`      //创建人
	UpdateBy   uint   `json:"update_by" gorm:"not null"`            //更新人
	ArticleId  uint   `json:"article_id" gorm:"not null"`           //文章/三月圈id
	CommentId  uint   `json:"comment_id" gorm:"not null;default:0"` //评论id
	AuthUserID uint `json:"auth_user_id" gorm:"not null;index"`   //作者id
	Type       uint8  `json:"type" gorm:"not null"`                 //消息类型 0文章 1三月圈
	Status     uint8  `json:"status" gorm:"not null"`               //消息状态 文章0发布 1评论 2文章：0发布、1评论、2收藏、3转载、4赞 三月圈：0发布、1评论、2，3，4三种赞
	IsRead     uint8  `json:"is_read" gorm:"not null;default:0"`    //0未读 1已读
	Title      string `json:"title" gorm:"not null;default:''"`     //文章题目
	Image      string `json:"image" gorm:"not null;default:''"`     //首图url
	Content    string `json:"content" gorm:"not null;default:''"`   //内容
}

func (e *Message) TableName() string {
	return `message`
}