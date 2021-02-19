package bo

type Article struct {
	ArticleMsg
	UserMsg
	FavourTotal  int64 `json:"favourTotal"`
	CollectTotal int64 `json:"collectTotal"`
	CommentTotal int64 `json:"commentTotal"`
	IsFollow     int   `json:"isFollow"`
}

type ArticleMsg struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Image      string `json:"image"`
	Tag        uint8  `json:"tag"`
	Kind       uint8  `json:"kind"`
	Type       uint   `json:"type"`
	UpdateBy   uint   `json:"update_by"`
	CreateBy   uint   `json:"create_by"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type UserMsg struct {
	UserID     int    `json:"user_id"`
	Nickname   string `json:"nickname" gorm:"column:nick_name"`
	AvatarPath string `json:"avatarPath" gorm:"column:avatar_path"`
}
