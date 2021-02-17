package models

type Follow struct {
	BaseModel
	CreateBy     uint `json:"create_by" gorm:"not null;index"`     //创建人
	UpdateBy     uint `json:"update_by" gorm:"not null"`            //更新人
	FollowUserId uint `json:"follow_user_id" gorm:"not null;index"` //关注人
}

func (e *Follow) TableName() string {
	return `follow`
}