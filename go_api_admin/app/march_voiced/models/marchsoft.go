package models

type MarchSoft struct {
	BaseModel
	CreateBy uint   `json:"create_by" gorm:"not null;index"` //创建人
	UpdateBy uint   `json:"update_by" gorm:"not null"`       //更新人
	Image    string `json:"image" gorm:"not null;default:''"`           //首图url
	Content  string `json:"content" gorm:"type(text);not null"`         //内容
}

func (e *MarchSoft) TableName() string {
	return `marchsoft`
}
