package bo

type GetFollowList struct {
	Follow []FollowInfo
}

type FollowInfo struct {
	FansTotal    int64  `json:"fansTotal"`   // 粉丝数
	FollowTotal  int64  `json:"followTotal"` // 关注数
	Signature    string `json:"signature"`
	ArticleTotal int64  `json:"articleTotal"` // 文章数
	IsFollow     int    `json:"isFollow"`     // 是否关注
	User
}

// 表示一个人的信息
type User struct {
	Id         uint   `json:"id"`
	AvatarPath string `json:"avatarPath"`
	Nickname   string `json:"nickname" gorm:"column:nick_name"`
}
