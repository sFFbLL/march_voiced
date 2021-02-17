package models

type ArticleComment struct {
	ArticleId uint `json:"article_id" gorm:""`
	Pid       uint `json:"pid" gorm:""`
	ReplyId   uint `json:"reply_id" gorm:""`
	CreateBy  uint `json:"creat_by" gorm:""`
	UpdateBy  uint `json:"update_by" gorm:""`
	BaseModel
}

func (a *ArticleComment) TableName() string {
	return `article_comment`
}
