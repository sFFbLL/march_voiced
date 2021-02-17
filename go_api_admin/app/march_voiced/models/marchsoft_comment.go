package models

type MarchSoftComment struct {
	BaseModel
	CreateBy    uint `json:"create_by" gorm:"not null;index"`    //创建人
	UpdateBy    uint `json:"update_by" gorm:"not null"`          //更新人
	MarchsoftId uint `json:"marchsoft_id" gorm:"not null;index"` //三月圈id
	Pid         uint `json:"pid" gorm:"not null"`                //上级评论id
	ReplyId     uint `json:"reply_id" gorm:"not null;index"`     //回复人id
}

func (e *MarchSoftComment) TableName() string {
	return `marchsoft_comment`
}
