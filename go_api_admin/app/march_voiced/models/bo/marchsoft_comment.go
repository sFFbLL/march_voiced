package bo

type GetMarchsoftComment struct {
	CommentSum []SignalMarchsoftComment
}

type SignalMarchsoftComment struct {
	MarchsoftComment
	ChildComments []MarchsoftComment
}

type ChildrenMarchsoftComment struct {
	Comment []MarchsoftComment
}

type MarchsoftComment struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"createTime"`
	MarchsoftCreate
	MarchsoftReply
}

type MarchsoftCreate struct {
	CreateBy     uint   `json:"createBy"`
	CreateByName string `json:"createByName"`
	IdAvatar     string `json:"idAvatar"`
}

type MarchsoftReply struct {
	ReplyId   uint   `json:"replyId"`
	ReplyName string `json:"replyName"`
}
