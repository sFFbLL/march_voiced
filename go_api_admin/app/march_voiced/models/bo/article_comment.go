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
	ArticleCreate
	ArticleReply
}

type ArticleCreate struct {
	CreateBy     uint   `json:"createBy"`
	CreateByName string `json:"createByName"`
	IdAvatar     string `json:"id_avatar"`
}

type ArticleReply struct {
	ReplyId   uint   `json:"replyId"`
	ReplyName string `json:"replyName"`
}
