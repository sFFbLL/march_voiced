package bo

type GetMessage struct {
	Current uint              `json:"current"` //页码
	Size    uint              `json:"size"`    //页内大小
	Pages   int               `json:"pages"`   //总页数
	Total   int64             `json:"total"`   //总条数
	Records *[]GetMessageData `json:"records"` //消息
}

type GetMessageData struct {
	Type         uint8  `json:"type"`                                 //消息类型（0 文章 / 1 三月圈）
	Status       uint8  `json:"status"`                               //消息状态 文章：0发布、1评论、2收藏、3转 载、4赞 三月圈：0发布、1评论、2，3，4三种赞 不返回收藏和评论
	ArticleId    uint   `json:"articleId" gorm:"column:article_id"`   //文章id/三月圈id
	UserId       uint   `json:"userId" gorm:"column:create_by"`       //用户id
	CreateTime   int64  `json:"createTime"`                           //创建时间
	Title        string `json:"title"`                                //标题
	Comment      string `json:"comment"`                              //评论
	Content      string `json:"content"`                              //内容
	Image        string `json:"image"`                                //首图url
	AvatarPath   string `json:"avatarPath" gorm:"column:avatar_path"` //头像
	Nickname     string `json:"nickname" gorm:"column:nick_name"`     //昵称
	ArticleTotal
}

type MarchSoftInfo struct {
	Title        string `json:"title"`
	Brief        string `json:"brief"`
	PersonTotal  uint   `json:"personTotal"`
	ArticleTotal uint   `json:"articleTotal"`
	IsMarch      uint8  `json:"is_march"`
}

type March struct {
	MarchMsg
	UserMsg
	MarchTotal
}

type MarchByUserId struct {
	MarchMsg
	MarchTotal
}

type MarchMsg struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	Image      string `json:"image"`
	UpdateBy   uint   `json:"update_by"`
	CreateBy   uint   `json:"create_by"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type MarchTotal struct {
	FaceTotal    int      `json:"faceTotal"`
	LikeTotal    int      `json:"likeTotal"`
	FavourTotal  int      `json:"favourTotal"`
	Type         int      `json:"type"`
	ImageList    []string `json:"imageList"`
	CommentTotal int64    `json:"commentTotal"`
	IsFollow     int      `json:"isFollow"`
}

type GoMarchMsg struct {
	MarchTotal
	MarchId     uint
	UserId      uint
	MarchUserId uint
}
