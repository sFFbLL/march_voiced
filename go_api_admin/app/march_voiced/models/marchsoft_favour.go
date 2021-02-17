package models

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
