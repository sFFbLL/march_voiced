package models

import (
	"project/common/global"

	"errors"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"

	"go.uber.org/zap"
)

type Follow struct {
	BaseModel
	CreateBy uint `json:"create_by" gorm:"not null;index"` //创建人
	UpdateBy uint `json:"update_by" gorm:"not null"`       //更新人
	FollowId uint `json:"follow_id" gorm:"not null;index"` //关注人
}

func (fo *Follow) TableName() string {
	return `follow`
}

// 判断userId 是否关注了 followId
func (fo *Follow) IsFollow(userId, followId int) (int, error) {
	var count int64
	var err error
	err = global.Eloquent.Table(fo.TableName()).Where("follow_id = ? AND create_by = ? AND is_deleted = 0", followId, userId).Count(&count).Error
	if err != nil {
		zap.L().Error("IsFollow Select failed", zap.Error(err))
		return 0, err
	}
	if count > 0 {
		return 1, err
	}
	return 0, err
}

// GetFollowList 查询关注列表信息的数据持久层
func (fo *Follow) GetFollowList(p *dto.GetFollowList) (followList *[]bo.FollowInfo, err error) {
	followList = new([]bo.FollowInfo)
	err = global.Eloquent.Table("sys_user").Joins("left join follow on sys_user.id = follow.follow_id").
		Where("follow.is_deleted=? AND follow.create_by=?", []byte{0}, p.Id).
		Offset(int((p.Current - 1) * p.Size)).Limit(int(p.Size)).Scan(followList).Error
	if err != nil {
		return nil, err
	}
	return followList, nil
}

// GetFollowList 查询粉丝列表信息的数据持久层
func (fo *Follow) GetFansList(p *dto.GetFollowList) (followList *[]bo.FansInfo, err error) {
	followList = new([]bo.FansInfo)
	err = global.Eloquent.Table("sys_user").
		Joins("left join follow on sys_user.id = follow.create_by").
		Where("follow.follow_id=? AND follow.is_deleted=?", p.Id, []byte{0}).
		Offset(int((p.Current - 1) * p.Size)).Limit(int(p.Size)).Scan(followList).Error
	if err != nil {
		return nil, err
	}
	return followList, nil
}

// GetFansTotal 查询一个人的粉丝总人数
func (fo *Follow) GetFansTotal(id uint) (fansTotal int64, err error) {
	err = global.Eloquent.Table(fo.TableName()).Where("is_deleted=? AND follow_id=?", []byte{0}, id).Count(&fansTotal).Error
	if err != nil {
		return 0, err
	}
	return fansTotal, nil
}

// GetFollowTotal 查询一个人的关注总人数
func (fo *Follow) GetFollowTotal(id uint) (followTotal int64, err error) {
	err = global.Eloquent.Table(fo.TableName()).Where("is_deleted=? AND create_by=?", []byte{0}, id).Count(&followTotal).Error
	if err != nil {
		return 0, err
	}
	return followTotal, nil
}

// GetFollowStatus 查询某人的关注状态
func (fo *Follow) GetFollowStatus(id int, me int) (err error) {
	var num int64
	err = global.Eloquent.Table(fo.TableName()).Where("is_deleted=? AND create_by=? AND follow_id=?", []byte{0}, me, id).Count(&num).Error
	if err != nil || num == 0 {
		return errors.New("未关注")
	}
	return nil
}

func (fo *Follow) UpdateStatus(id int, me int, status int) (err error) {
	// UpdateStatus 修改关注状态

	if status == 2 {
		err = global.Eloquent.Table(fo.TableName()).Where("is_deleted=? AND create_by=? AND follow_id=?", 0, me, id).Update("is_deleted", 1).Error
		if err != nil {
			return errors.New("修改失败")
		}
		return nil
	}
	err = fo.GetFollowStatus(id, me)
	if err != nil {
		fo.CreateBy = uint(me)
		fo.UpdateBy = uint(me)
		fo.FollowId = uint(id)
		err = global.Eloquent.Table(fo.TableName()).Create(fo).Error
		if err != nil {
			return errors.New("修改失败")
		}
		go AddFollowMessage(uint(me), uint(id))
	}
	return nil
}
