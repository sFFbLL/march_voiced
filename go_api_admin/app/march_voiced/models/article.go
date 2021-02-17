package models

type Article struct {
	Title       string `json:"title" gorm:"size:128;"`
	Content     string `json:"content" gorm:"type(text)"`
	Image       string `json:"image" gorm:"size:128;"`
	Tag         uint8  `json:"tag" gorm:"size:5;"`
	Kind        uint8  `json:"kind" gorm:"size:4;"`
	Status      uint8  `json:"status" gorm:"size:2;"`
	Type        uint8  `json:"type" gorm:"size:1;"`
	IsRecommend uint8  `json:"is_recommend" gotm:"size:1;DEFAULT:0;"`
	CreatBy     uint   `json:"creat_by" gorm:""`
	UpdateBy    uint   `json:"update_by" gorm:""`
	BaseModel
}

func (a *Article) TableName() string {
	return `article`
}
