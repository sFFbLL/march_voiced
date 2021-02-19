package bo

type GetFollowList struct {
	Follow []FollowInfo
}

type FollowInfo struct {
	FansTotal   int64  `json:"fans_total"`   // 粉丝数
	FollowTotal int64  `json:"follow_total"` // 关注数
	Signature   string `json:"signature"`
	User
}

// 表示一个人的信息
type User struct {
	Id         int    `json:"id"`
	AvatarPath string `json:"avatar_path"`
	Nickname   string `json:"nickname" gorm:"column:nick_name"`
}
