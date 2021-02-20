package dto

type InsertArticleDto struct {
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Image     string `json:"image"`
	Kind      uint8  `json:"kind" binding:"required"`
	Tag       uint8  `json:"tag" binding:"required"`
	Status    *uint8 `json:"status" binding:"required"`
	Type      *uint  `json:"type" binding:"required"`
	WordCount uint   `json:"word_count" binding:"required"`
}

type UpdateArticleDto struct {
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Image     string `json:"image"`
	ID        uint   `json:"id"  binding:"required"`
	WordCount uint   `json:"word_count" binding:"required"`
	Type      *uint  `json:"type"`
	Status    *uint8 `json:"status" binding:"required,lte=1"`
	Kind      uint8  `json:"kind"`
	Tag       uint8  `json:"tag"`
}

type Paging struct {
	Current int `form:"current"`
	Size    int `form:"size"`
}

type SelectArticleByUser struct {
	Paging
	ID   uint  `form:"id" binding:"required"`
	Kind uint8 `form:"kind" binding:"required,lte=2"`
}

type ArticlePass struct {
	Id     uint   `json:"id" binding:"required"`
	Status *uint8 `json:"status" binding:"required,lte=1"`
}
