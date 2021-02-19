package dto

type GetFollowList struct {
	Id      int `form:"id"` // 用户id
	Current int `form:"current"`
	Size    int `form:"size"`
}
