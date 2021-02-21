package models

import "project/common/global"

type MarchSoftInfo struct {
	ID           int    `json:"id" gorm:"primary_key"` //ID
	Title        string `json:"title"`
	Brief        string `json:"brief"`
	PersonTotal  uint   `json:"person_total"`
	ArticleTotal uint   `json:"article_total"`
}

func (m *MarchSoftInfo) TableName() string {
	return "marchsoft_info"
}

func (m *MarchSoftInfo) GetMarchSoftInfo() error {
	return global.Eloquent.Table(m.TableName()).Where("id = 1").Find(m).Error
}
