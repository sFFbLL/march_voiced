package dto

type AddMarchsoftComment struct {
	ID      uint   `json:"id"`       // 文章id
	ReplyId uint   `json:"reply_id"` // 回复人id
	Pid     uint   `json:"pid"`      // 父评论id
	Content string `json:"content"`  // 评论内容
}

type GetMarchsoftComment struct {
	ID        uint   `form:"id" binding:"required"` // 文章id
	Current   uint   `form:"current"`               // 当前页
	Size      uint   `form:"size"`
	ChildSize uint   `form:"childSize"` // 多少条子评论
	Orders    string `form:"orders"`    // 排序规则
}

type GetMarchsoftChildComment struct {
	ID      uint   `form:"id" binding:"required"`      // 一级评论的id
	Current uint   `form:"current" binding:"required"` // 当前页
	Size    uint   `form:"size" `
	Orders  string `form:"orders"` // 排序规则
}
