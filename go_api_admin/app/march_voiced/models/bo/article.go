package bo

type Article struct {
	ArticleMsg
	UserMsg
	ArticleTotal
}

type ArticleMsg struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Describe   string `json:"describe"`
	Image      string `json:"image"`
	Status     uint8  `json:"status"`
	Tag        uint   `json:"tag_id"`
	Type       uint   `json:"type"`
	UpdateBy   uint   `json:"update_by"`
	CreateBy   uint   `json:"create_by"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type ArticleDetail struct {
	ArticleDetailMsg
	UserMsg
	ArticleTotal
	IsCollect int `json:"isCollect"`
	IsFavour  int `json:"isFavour"`
}

type ArticleDetailMsg struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Image      string `json:"image"`
	Tag        string `json:"tag"`
	Kind       uint8  `json:"kind"`
	Status     uint8  `json:"status"`
	TagId      uint   `json:"tag_id"`
	WordCount  uint   `json:"word_count"`
	Type       uint   `json:"type"`
	UpdateBy   uint   `json:"update_by"`
	CreateBy   uint   `json:"create_by"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type UserMsg struct {
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
	MessageId     uint
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
	Icon        string `json:"icon"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

type ApplyArticleList struct {
	Current uint                    `json:"current"` //页码
	Size    uint                    `json:"size"`    //页内大小
	Pages   int                     `json:"pages"`   //总页数
	Total   int64                   `json:"total"`   //总条数
	Records *[]ApplyArticleListData `json:"records"` //消息
}

type ApplyArticleListData struct {
	Title       string `json:"title"`
	Tag         string `json:"tag"`
	Nickname    string `json:"nickname" gorm:"column:nick_name"`
	Id          uint   `json:"id"`
	CreateTime  int64  `json:"createTime" gorm:"column:status_update_time"`
	Status      uint8  `json:"status"`
	IsRecommend uint8  `json:"recommend" gorm:"column:is_recommend"`
}

type ArticleCollectByUserId struct {
	Current uint       `json:"current"` //页码
	Size    uint       `json:"size"`    //页内大小
	Pages   int        `json:"pages"`   //总页数
	Total   int64      `json:"total"`   //总条数
	Records *[]Article `json:"records"` //消息
}
