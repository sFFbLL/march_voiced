package bo

type GetArticleComment struct {
	CommentSum []SignalArticleComment
}

type SignalArticleComment struct {
	ArticleComment
	ChildComments []ArticleComment
}

type ChildrenComment struct {
	Comment []ArticleComment
}

type ArticleComment struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	ArticleCreate
	ArticleReply
}

type ArticleCreate struct {
	CreateBy     uint   `json:"create_by"`
	CreateByName string `json:"create_by_name"`
	IdAvatar     string `json:"id_avatar"`
}

type ArticleReply struct {
	ReplyId   uint   `json:"reply_id"`
	ReplyName string `json:"reply_name"`
}
