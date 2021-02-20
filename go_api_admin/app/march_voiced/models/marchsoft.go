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

func (e *MarchSoft) GetMarchApplyUser(applyMarchUser *bo.ApplyMarchUser, p *dto.ApplyMarchPaginator) (err error) {
	nickname :="%" + p.Nickname + "%"
	table := global.Eloquent.Table("sys_user").
		Select("sys_user.id, sys_user.nick_name, sys_user.avatar_path, sys_user.phone, sys_user.march_update_time, sys_user.is_march, sys_dept.name").
		Joins("left join sys_dept on sys_user.dept_id = sys_dept.id").
		Where("sys_user.is_deleted=0 and sys_user.is_march=2 and sys_user.nickname like ?", nickname)
	if p.EndTime != 0 && p.StartTime != 0 {
		err = table.Where("sys_user.march_update_time > ? AND sys_user.march_update_time < ?", p.StartTime, p.EndTime).Count(&applyMarchUser.Total).
			Order("sys_user.march_update_time desc").Limit(int(p.Size)).Offset(int(p.Current - 1*p.Size)).
			Find(applyMarchUser.Records).Error
	} else {
		err = table.Count(&applyMarchUser.Total).
			Order("sys_user.march_update_time desc").Limit(int(p.Size)).Offset(int(p.Current - 1*p.Size)).
			Find(applyMarchUser.Records).Error
	}
	return
}
