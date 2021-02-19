package dto

type MarchPass struct {
	Id     uint   `json:"id" binding:"required"`
	Status *uint8 `json:"status" binding:"required,lte=1"`
}
