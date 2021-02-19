package service

import (
	models2 "project/app/admin/models"
	"project/app/march_voiced/models"
	"project/common/global"
)

type Marchsoft struct {
}

func (e *Marchsoft) ApplyMarch(userId int) (err error) {
	err = global.Eloquent.Table("sys_user").
		Where("id=? and is_deleted=0", userId).
		Updates(&models2.SysUser{IsMarch: 1}).Error
	if err != nil {
		return
	}

	go models.AddSysMessage(1, 1, uint(userId))
	return
}
