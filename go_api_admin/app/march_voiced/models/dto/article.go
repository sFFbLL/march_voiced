package dto

type InsertArticleDto struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Image   string `json:"image"`
	Kind    uint8  `json:"kind" binding:"required"`
	Tag     uint8  `json:"tag" binding:"required"`
	Status  *uint8 `json:"status" binding:"required"`
	Type    *uint8 `json:"type" binding:"required"`
}

type UpdateArticleDto struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Image   string `json:"image"`
	ID      uint   `json:"id"  binding:"required"`
	Kind    uint8  `json:"kind"`
	Tag     uint8  `json:"tag"`
	Status  *uint8 `json:"status"  binding:"required"`
	Type    *uint8 `json:"type"`
}
