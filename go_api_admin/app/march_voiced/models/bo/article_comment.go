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
	CreateTime int64  `json:"createTime"`
	ChildSum   int64  `json:"childSum"`
	ArticleCreate
	ArticleReply
}

type ArticleCreate struct {
	CreateBy     uint   `json:"createBy"`
	CreateByName string `json:"createByName"`
	IdAvatar     string `json:"idAvatar"`
}

type ArticleReply struct {
	ReplyId   uint   `json:"replyId"`
	ReplyName string `json:"replyName"`
}
