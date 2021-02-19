package dto

type AddMarchsoftComment struct {
	ID      uint   `json:"id"`       // 文章id
	Content string `json:"content"`  // 评论内容
	ReplyId uint   `json:"reply_id"` // 回复人id
	Pid     uint   `json:"pid"`      // 父评论id
}

type DeleteMarchsoftComment struct {
	ID int `json:"id"`
}

type GetMarchsoftComment struct {
	ID      int    `form:"id" binding:"required"`      // 文章id
	Current int    `form:"current" binding:"required"` // 当前页
	Size    int    `form:"size" binding:"required"`
	Orders  string `form:"orders"` // 排序规则
}

type GetMarchsoftChildComment struct {
	ID      int    `form:"id" binding:"required"`      // 一级评论的id
	Current int    `form:"current" binding:"required"` // 当前页
	Size    int    `form:"size" `
	Orders  string `form:"orders"` // 排序规则
}
