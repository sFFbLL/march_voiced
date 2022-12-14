package dto

type AddArticleComment struct {
	ID       uint   `json:"id"`       // 文章id
	ReplyId  uint   `json:"replyId"`  // 回复人id
	FollowId uint   `json:"followId"` // 父评论id
	Content  string `json:"content"`  // 评论内容
}

type GetArticleComment struct {
	ID        uint   `form:"id" binding:"required"` // 文章id
	Current   uint   `form:"current"`               // 当前页
	Size      uint   `form:"size"`
	ChildSize uint   `form:"childSize"` // 多少条子评论
	Orders    string `form:"orders"`    // 排序规则
}

type GetArticleChildComment struct {
	ID      uint   `form:"id" binding:"required"`      // 一级评论的id
	Current uint   `form:"current" binding:"required"` // 当前页
	Size    uint   `form:"size" binding:"required"`
	Orders  string // 排序规则
}
