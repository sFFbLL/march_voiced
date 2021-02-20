package service

import (
	"errors"
	models2 "project/app/admin/models"
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"
)

type Marchsoft struct {
}

// GetMarchApplyUser 业务方法
func (e *Marchsoft) GetMarchApplyUser(p *dto.Paginator, userId int) (applyMarchUser *bo.ApplyMarchUser, err error) {
	marchSoft := new(models.MarchSoft)
	applyMarchUser = new(bo.ApplyMarchUser)
	applyMarchUserMessage := new([]bo.ApplyMarchUserMessage)
	applyMarchUser.Message = applyMarchUserMessage
	err = marchSoft.GetMarchApplyUser(applyMarchUser, p)
	return
}

// ApplyMarch 申请加入三月圈业务方法
func (e *Marchsoft) ApplyMarch(userId int) (err error) {
	var count int64
	table := global.Eloquent.Table("sys_user").
		Where("id=? and is_march=0 and is_deleted=0", userId)
	err = table.Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		return errors.New("用户已申请或已加入")
	}
	var i uint8 = 2
	err = table.Updates(&models2.SysUser{IsMarch: &i}).Error
	if err != nil {
		return
	}

	go models.AddSysMessage(1, i, uint(userId), uint(userId))
	return
}

// MarchPass （后台）三月圈审核通过或驳回业务方法
func (e *Marchsoft) MarchPass(p *dto.MarchPass, userId int) (err error) {
	var count int64
	table := global.Eloquent.Table("sys_user").
		Where("id=? and is_march=2 and is_deleted=0", p.Id)
	err = table.Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		return errors.New("该用户未申请三月圈")
	}
	err = table.Updates(&models2.SysUser{IsMarch: p.Status}).Error
	if err != nil {
		return
	}
	go models.AddSysMessage(1, *p.Status, uint(userId), p.Id)
	return
}
