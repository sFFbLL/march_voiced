package bo

type Article struct {
	ID           int     `json:"id"`
	Title        string  `json:"title" `
	Content      string  `json:"content"`
	Image        string  `json:"image"`
	Tag          uint8   `json:"tag"`
	Kind         uint8   `json:"kind"`
	Type         uint    `json:"type"`
	FavourTotal  int64   `json:"favourTotal"`
	CollectTotal int64   `json:"collectTotal"`
	CommentTotal int64   `json:"commentTotal"`
	CreatTime    int64   `json:"creat_time"`
	UpdateTime   int64   `json:"update_time"`
	CreatBy      uint    `json:"creat_by"`
	UpdateBy     uint    `json:"update_by"`
	User         UserMsg `json:"user"`
}

type UserMsg struct {
	ID         int    `json:"id"`
	IsFollow   int    `json:"isFollow"`
	Nickname   string `json:"nickname" gorm:"column:nick_name"`
	AvatarPath string `json:"avatarPath"`
}
