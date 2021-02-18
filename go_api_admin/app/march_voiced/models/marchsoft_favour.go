package models

import (
	"errors"

	"project/app/march_voiced/models/dto"
	"project/common/global"
)

type MarchSoftFavour struct {
	BaseModel
	CreateBy    uint  `json:"create_by" gorm:"not null;index"`    //创建人
	UpdateBy    uint  `json:"update_by" gorm:"not null"`          //更新人
	MarchsoftId uint  `json:"marchsoft_id" gorm:"not null;index"` //三月圈id
	Type        uint8 `json:"type" gorm:"not null"`               //赞类型 0表情 1喜欢 2散花
}

func (e *MarchSoftFavour) TableName() string {
	return `marchsoft_favour`
}

func (e *MarchSoftFavour) AddMarchFavour(p *dto.MarchFavourDto) (err error) {
	tx := global.Eloquent.Table(e.TableName()).Begin()
	var count int64
	err = tx.Where("create_by=? and marchsoft_id=? and type=? and is_deleted=?", e.CreateBy, e.MarchsoftId, e.Type, 0).Count(&count).Error
	if err != nil {
		return errors.New("查询操作失败")
	}
	// 先删除数据
	err = tx.Where("create_by=? and marchsoft_id=? and is_deleted=?", e.CreateBy, e.MarchsoftId, 0).Update("is_deleted", 1).Error
	// 如果有就返回
	if err != nil {
		tx.Rollback()
		return
	}
	if count>0 {
		tx.Commit()
		return
	}
	// 如果没有就创建
	err = tx.Create(e).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (e *MarchSoftFavour) AddMarchFavourMessage(s *uint8, userId uint) {
	var count int64
	err := global.Eloquent.Table(e.TableName()).Where("create_by=? and marchsoft_id=? and is_deleted=1", e.CreateBy, e.MarchsoftId).Count(&count).Error
	if err != nil || count > 0 {
		return
	}
	AddMessage(0, *s+2, e.MarchsoftId, userId, "")
}

func (e *MarchSoftFavour) MarchFavourCount() (count int64, err error) {
	err = global.Eloquent.Table(e.TableName()).
		Where("article_id=? and is_deleted=0", e.MarchsoftId).Count(&count).Error
	return
}
