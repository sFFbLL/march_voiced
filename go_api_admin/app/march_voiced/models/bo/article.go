package bo

type Article struct {
	ArticleMsg
	UserMsg
	ArticleTotal
}

type ArticleMsg struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Image      string `json:"image"`
	Tag        uint8  `json:"tag"`
	Kind       uint8  `json:"kind"`
	WordCount  uint   `json:"word_count"`
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

type WordRes struct {
	Sensitive []SensitiveWord
}

type SensitiveWord struct {
	Word    string `json:"word"`
	Indexes []int  `json:"indexes"`
	Length  int    `json:"length"`
}

type ArticleTotal struct {
	FavourTotal  int64 `json:"favourTotal"`
	CollectTotal int64 `json:"collectTotal"`
	CommentTotal int64 `json:"commentTotal"`
	IsFollow     int   `json:"isFollow"`
}

type GoArticleMsg struct {
	ArticleTotal
	UserId        uint
	ArticleId     uint
	ArticleUserId uint
	CreateTime    int64
}

type ArticleUser struct {
	ArticleMsg
	ArticleTotal
}

type IsFavourCollectByArticleId struct {
	IsFavour  int `json:"isFavour"`
	IsCollect int `json:"isCollect"`
}

type ArticleTagList struct {
	ID          int    `json:"id"`
	Icon        int    `json:"icon"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}
