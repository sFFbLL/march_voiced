package bo

type GetFollowList struct {
	Follow []FollowInfo `json:"follow"`
}

type GetFansList struct {
	Fans []FansInfo `json:"follow"`
}

type FollowInfo struct {
	User
	FollowId     int64  `json:"id" gorm:"follow_id"`
	FansTotal    int64  `json:"fansTotal"`   // 粉丝数
	FollowTotal  int64  `json:"followTotal"` // 关注数
	IsFollow     int    `json:"isFollow"`    // 是否关注
	Signature    string `json:"signature"`
	ArticleTotal int64  `json:"articleTotal"` // 文章数
}

type FansInfo struct {
	User
	FansId       int64  `json:"id" gorm:"column:create_by"`
	FansTotal    int64  `json:"fansTotal"`   // 粉丝数
	FollowTotal  int64  `json:"followTotal"` // 关注数
	IsFollow     int    `json:"isFollow"`    // 是否关注
	Signature    string `json:"signature"`
	ArticleTotal int64  `json:"articleTotal"` // 文章数
}

// 表示一个人的信息
type User struct {
	AvatarPath string `json:"avatarPath"`
	Nickname   string `json:"nickname" gorm:"column:nick_name"`
}
