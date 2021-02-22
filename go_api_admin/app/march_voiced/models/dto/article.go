package dto

type ApplyArticlePaginator struct {
	Current   uint   `form:"current"`
	Size      uint   `form:"size"`
	EndTime   uint   `form:"endTime"`   //结束时间
	StartTime uint   `form:"startTime"` //创建时间
	Status    uint8  `form:"status"`    //1通过 2审核中
	Nickname  string `form:"nickname"`
	Title     string `form:"title"`
}

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
	Size    int `form:"size" binding:"lte=10"`
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
