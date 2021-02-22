package models

import "project/common/global"

type MarchSoftImage struct {
	BaseModel
	CreateBy    uint   `json:"create_by" gorm:"not null;index"`    //创建人
	UpdateBy    uint   `json:"update_by" gorm:"not null"`          //更新人
	MarchsoftId uint   `json:"marchsoft_id" gorm:"not null;index"` //三月圈id
	Image       string `json:"image" gorm:"not null"`              // 图片url
}

func (e *MarchSoftImage) TableName() string {
	return `marchsoft_image`
}

func (e *MarchSoftImage) GetImageListById() (imageList []string, err error) {
	err = global.Eloquent.Table(e.TableName()).Select("image").
		Where("marchsoft_id = ? AND is_deleted = 0 ", e.MarchsoftId).Find(&imageList).Error
	return
}
