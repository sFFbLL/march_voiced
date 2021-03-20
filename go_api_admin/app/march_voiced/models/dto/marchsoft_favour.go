package dto

// MarchFavourDto 三月圈点赞dto
type MarchFavourDto struct {
	Id   uint  `json:"id" binding:"required"`
	Type uint8 `json:"type" binding:"required,lte=3"`
}
