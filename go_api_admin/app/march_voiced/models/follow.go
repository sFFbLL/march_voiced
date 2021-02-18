package models

import (
	"project/common/global"

	"go.uber.org/zap"
)

type Follow struct {
	BaseModel
	CreateBy uint `json:"create_by" gorm:"not null;index"` //创建人
	UpdateBy uint `json:"update_by" gorm:"not null"`       //更新人
	FollowId uint `json:"follow_id" gorm:"not null;index"` //关注人
}

func (e *Follow) TableName() string {
	return `follow`
}

// 判断userId 是否关注了 followId
func (e *Follow) IsFollow(userId, followId int) (bool, error) {
	var count int64
	var err error
	err = global.Eloquent.Table("follow").Where("follow_id = ? AND creat_by = ? AND is_deleted = 0", followId, userId).Count(&count).Error
	if err != nil {
		zap.L().Error("IsFollow Select failed", zap.Error(err))
		return false, err
	}
	if count > 0 {
		return true, err
	}
	return false, err
}
