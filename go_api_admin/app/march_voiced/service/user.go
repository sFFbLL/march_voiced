package service

import (
	"fmt"
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
)

type User struct {
}

func (u *User) SelectUserInfo(id int, me int) (res *bo.SelectUserInfo, err error) {
	user := new(models.User)
	fo := new(models.Follow)
	if id == 0 {
		id = me
	}
	err = user.SelectUserInfo(id)
	followTotal, err := fo.GetFollowTotal(id)
	fansTotal, err := fo.GetFansTotal(id)
	isFollow, err := fo.IsFollow(me, id)
	fmt.Println(isFollow)
	articleTotal := 65
	collectTotal, err := user.GetUserCollect(id)
	if err != nil {
		return
	}
	info := &bo.UserIn{
		Id:         user.ID,
		AvatarPath: user.AvatarPath,
		Nickname:   user.NickName,
		IsFollow:   isFollow,
	}
	res = &bo.SelectUserInfo{
		Signature:    user.Signature,
		FollowTotal:  followTotal,
		FansTotal:    fansTotal,
		ArticleTotal: articleTotal,
		CollectTotal: collectTotal,
		UserIn:       *info,
	}

	return
}
