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
	CreateTime int64  `json:"create_time"`
	ChildSum   int64  `json:"childSum"`
	MarchsoftCreate
	MarchsoftReply
}

type MarchsoftCreate struct {
	CreateBy     uint   `json:"createBy"`
	CreateByName string `json:"createByName"`
	IdAvatar     string `json:"id_avatar"`
}

type MarchsoftReply struct {
	ReplyId   uint   `json:"replyId"`
	ReplyName string `json:"replyName"`
}
