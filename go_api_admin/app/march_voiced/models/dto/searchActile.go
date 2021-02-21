package dto

// SearchActileDto 搜索文章
type SearchActileDto struct {
	SearchWord string `form:"searchWord" binding:"required"`
	Current    int `form:"current" binding:"required"`
	Size 		int `form:"size" binding:"required"`
}

//
type Article struct {
	Id int `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
	CreateBy int  `db:"create_by" json:"create_by"`
	UpdateBy int `db:"update_by" json:"update_by"`
	CreateTime int `db:"create_time" json:"create_time"`
	UpdateTime int `db:"update_time" json:"update_time"`
}