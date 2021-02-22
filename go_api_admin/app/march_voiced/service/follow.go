package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
)

type Follow struct {
}

// GetFollowList 查询某人的关注列表
func (fo *Follow) GetFollowList(p *dto.GetFollowList) (*bo.GetFollowList, error) {
	follow := new(models.Follow)
	var res bo.GetFollowList
	var signalFollowList []bo.FollowInfo
	if p.Size == 0 && p.Current == 0 {
		p.Size = 5
		p.Current = 1
	}
	followList, err := follow.GetFollowList(p)
	if err != nil {
		return nil, err
	}
	for _, followTwo := range *followList {
		fansTotal, err := follow.GetFansTotal(followTwo.Id)
		if err != nil {
			return nil, err
		}
		followTotal, err := follow.GetFollowTotal(followTwo.Id)
		if err != nil {
			return nil, err
		}
		followTwo.FansTotal = fansTotal
		followTwo.FollowTotal = followTotal
		signalFollowList = append(signalFollowList, followTwo)
	}
	res.Follow = signalFollowList
	return &res, nil
}

// GetFansList 查询粉丝列表
func (fo *Follow) GetFansList(p *dto.GetFollowList) (*bo.GetFollowList, error) {
	follow := new(models.Follow)
	var res bo.GetFollowList
	var signalFollowList []bo.FollowInfo
	if p.Size == 0 && p.Current == 0 {
		p.Size = 5
		p.Current = 1
	}
	followList, err := follow.GetFansList(p)
	if err != nil {
		return nil, err
	}
	for _, followTwo := range *followList {
		fansTotal, err := follow.GetFansTotal(followTwo.Id)
		if err != nil {
			return nil, err
		}
		followTotal, err := follow.GetFollowTotal(followTwo.Id)
		if err != nil {
			return nil, err
		}
		followTwo.FansTotal = fansTotal
		followTwo.FollowTotal = followTotal
		signalFollowList = append(signalFollowList, followTwo)
	}
	res.Follow = signalFollowList
	return &res, nil
}

// GetFollowStatus 得到关注状态
func (fo *Follow) GetFollowStatus(id int, me int) (err error) {
	follow := new(models.Follow)
	err = follow.GetFollowStatus(id, me)
	return
}

// UpdateStatus 更新关注状态
func (fo *Follow) UpdateStatus(id int, me int) (err error) {
	follow := new(models.Follow)
	err = follow.UpdateStatus(id, me)
	return
}
