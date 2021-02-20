package models

import (
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"
)

type MarchSoft struct {
	BaseModel
	CreateBy uint   `json:"create_by" gorm:"not null;index"`    //创建人
	UpdateBy uint   `json:"update_by" gorm:"not null"`          //更新人
	Image    string `json:"image" gorm:"not null;default:''"`   //首图url
	Content  string `json:"content" gorm:"type(text);not null"` //内容
}

func (e *MarchSoft) TableName() string {
	return `marchsoft`
}

func (e *MarchSoft) GetMarchApplyUser(applyMarchUser *bo.ApplyMarchUser, p *dto.Paginator) (err error) {
	err = global.Eloquent.Debug().Table("sys_message").Where("type=? and is_deleted=0", 1).
		Count(&applyMarchUser.Count).
		Order("create_time desc").Limit(int(p.Size)).Offset(int(p.Current - 1*p.Size)).
		Find(applyMarchUser.Message).Error
	return
}
