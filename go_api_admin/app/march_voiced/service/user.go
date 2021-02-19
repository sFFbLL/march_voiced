package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"

	"go.uber.org/zap"
)

type User struct {
}

func (u *User) SelectUserInfo(id int, me int) (res *bo.SelectUserInfo, err error) {
	user := new(models.User)
	fo := new(models.Follow)
	a := new(models.Article)
	if id == 0 {
		id = me
	}
	userIn, err := user.SelectUserInfo(id)
	if err != nil {
		zap.L().Error("SelectUserInfo failed", zap.Error(err))
		return
	}
	followTotal, err := fo.GetFollowTotal(id)
	if err != nil {
		zap.L().Error("GetFollowTotal failed", zap.Error(err))
		return
	}
	fansTotal, err := fo.GetFansTotal(id)
	if err != nil {
		zap.L().Error("GetFansTotal failed", zap.Error(err))
		return
	}
	isFollow, err := fo.IsFollow(me, id)
	if err != nil {
		zap.L().Error("IsFollow failed", zap.Error(err))
		return
	}
	articleTotal, err := a.ArticleCountByUserId(id)
	if err != nil {
		zap.L().Error("ArticleCountByUserId failed", zap.Error(err))
		return
	}
	collectTotal, err := user.GetUserCollect(id)
	if err != nil {
		zap.L().Error("GetUserCollect failed", zap.Error(err))
		return
	}
	info := &bo.UserIn{
		Id:         userIn.Id,
		AvatarPath: userIn.AvatarPath,
		Nickname:   userIn.NickName,
		IsFollow:   isFollow,
	}
	res = &bo.SelectUserInfo{
		Signature:    userIn.Signature,
		FollowTotal:  followTotal,
		FansTotal:    fansTotal,
		ArticleTotal: articleTotal,
		CollectTotal: collectTotal,
		UserIn:       *info,
	}
	return
}
