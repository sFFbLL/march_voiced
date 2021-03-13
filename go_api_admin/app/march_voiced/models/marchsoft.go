package models

import (
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"

	"go.uber.org/zap"
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
	nickname := "%" + p.Nickname + "%"
	table := global.Eloquent.Table("sys_user").
		Select("sys_user.id, sys_user.nick_name, sys_user.avatar_path, sys_user.phone, sys_user.march_update_time, sys_user.is_march, sys_dept.name").
		Joins("left join sys_dept on sys_user.dept_id = sys_dept.id").
		Where("sys_user.is_deleted=0 and sys_user.is_march=2 and sys_user.nick_name like ?", nickname)
	if p.EndTime != 0 && p.StartTime != 0 {
		err = table.Where("sys_user.march_update_time > ? AND sys_user.march_update_time < ?", p.StartTime, p.EndTime).Count(&applyMarchUser.Total).
			Order("sys_user.march_update_time desc").Limit(int(p.Size)).Offset(int((p.Current - 1) * p.Size)).
			Find(applyMarchUser.Records).Error
	} else {
		err = table.Count(&applyMarchUser.Total).
			Order("sys_user.march_update_time desc").Limit(int(p.Size)).Offset(int((p.Current - 1) * p.Size)).
			Find(applyMarchUser.Records).Error
	}
	return
}

func (e *MarchSoft) InsertMarchSoft(imgList []string) (err error) {
	var marchSoftImage MarchSoftImage
	marchSoftImageList := new([]MarchSoftImage)

	// 开启事务
	tx := global.Eloquent.Begin()
	err = tx.Table(e.TableName()).Create(e).Error
	if err != nil {
		zap.L().Error("InsertMarchSoft Dao Creat Failed", zap.Error(err))
		tx.Rollback()
		return
	}

	for i := 0; i < len(imgList); i++ {
		marchSoftImage = MarchSoftImage{
			MarchsoftId: uint(e.ID),
			Image:       imgList[i],
			CreateBy:    e.CreateBy,
			UpdateBy:    e.UpdateBy,
		}
		*marchSoftImageList = append(*marchSoftImageList, marchSoftImage)
	}
	if len(*marchSoftImageList) > 0 {
		err = tx.Table(marchSoftImage.TableName()).Create(marchSoftImageList).Error
		if err != nil {
			zap.L().Error("InsertMarchSoft Insert MarchSoftImages Failed", zap.Error(err))
			tx.Rollback()
			return
		}
	}

	err = tx.Commit().Error
	if err != nil {
		zap.L().Error("InsertMarchSoft Commit Failed", zap.Error(err))
	}
	return
}

func (e *MarchSoft) DeleteMarchSoft() (err error) {
	err = global.Eloquent.Table(e.TableName()).Where("id = ? and is_deleted = 0", e.ID).Updates(e).Error
	return
}

func (e *MarchSoft) SelectMarchSoftList(paging dto.Paging) (marchList *[]bo.March, err error) {
	marchList = new([]bo.March)
	err = global.Eloquent.Table(e.TableName()).
		Select("marchsoft.id, marchsoft.content, marchsoft.image, marchsoft.create_time, marchsoft.create_by, marchsoft.update_by, marchsoft.update_time, sys_user.nick_name, sys_user.avatar_path").
		Joins("JOIN sys_user ON marchsoft.create_by = sys_user.id").
		Where("marchsoft.is_deleted = 0").
		Limit(int(paging.Size)).Offset(int((paging.Current - 1) * paging.Size)).Find(marchList).Error
	return
}

func (e *MarchSoft) SelectMarchSoftListByUserId(paging dto.SelectMarchListById) (marchList *[]MarchSoft, err error) {
	marchList = new([]MarchSoft)
	err = global.Eloquent.Table(e.TableName()).
		Where("create_by = ? AND is_deleted = 0", paging.ID).
		Limit(int(paging.Size)).Offset(int((paging.Current - 1) * paging.Size)).Find(marchList).Error
	return
}

func (e *MarchSoft) MarchDetail() (march *bo.March, err error) {
	march = new(bo.March)
	err = global.Eloquent.Table(e.TableName()).
		Select("marchsoft.id, marchsoft.content, marchsoft.image , marchsoft.create_time, marchsoft.create_by, marchsoft.update_by, marchsoft.update_time, sys_user.nick_name, sys_user.avatar_path").
		Joins("JOIN sys_user ON marchsoft.create_by = sys_user.id").
		Where("marchsoft.id = ? AND marchsoft.is_deleted = 0", e.ID).First(march).Error
	return
}

func (e *MarchSoft) GetMarchById() (err error) {
	err = global.Eloquent.Table(e.TableName()).Where("id = ? AND is_deleted = 0", e.ID).First(e).Error
	return
}
