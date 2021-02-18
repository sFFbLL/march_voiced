package bo

type GetMarchsoftComment struct {
	CommentSum []SignalMarchsoftComment
}

type SignalMarchsoftComment struct {
	MarchsoftComment
	ChildComments []MarchsoftComment
}

type ChildrenArticleComment struct {
	Comment []MarchsoftComment
}

type MarchsoftComment struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	MarchsoftCreate
	MarchsoftReply
}

type MarchsoftCreate struct {
	CreateBy     int    `json:"create_by"`
	CreateByName string `json:"create_by_name"`
	IdAvatar     string `json:"id_avatar"`
}

type MarchsoftReply struct {
	ReplyId     int    `json:"reply_id"`
	ReplyName   string `json:"reply_name"`
	ReplyAvatar string `json:"reply_avatar"`
}