package dto

type GetFollowList struct {
	Id      uint `form:"id"` // 用户id
	Current uint `form:"current"`
	Size    uint `form:"size"`
}

type UpdateStatus struct {
	Id uint `json:"id"`
}
