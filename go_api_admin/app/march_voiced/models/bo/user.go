package bo

type SelectUserInfo struct {
	Signature    string `json:"signature"`
	FollowTotal  int64  `json:"followTotal"`
	FansTotal    int64  `json:"fansTotal"`
	ArticleTotal int64  `json:"articleTotal"`
	CollectTotal int64  `json:"collectTotal"`
	CraftTotal   int64  `json:"craftTotal"`
	IsMe         int    `json:"isMe"`
	UserIn
}

type UserIn struct {
	Id         int    `json:"id"`
	AvatarPath string `json:"avatarPath"`
	Nickname   string `json:"nickname"`
	IsFollow   int    `json:"isFollow"`
}
